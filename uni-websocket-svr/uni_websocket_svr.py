import json
import time
import random
import asyncio
import websockets
from game import GameRoom
from utils import orderList
from message import GameMessage

game_user_dict = { }

# 发送消息
async def sendMessage(websocket, msg):
    try :
        if websocket :
            await websocket.send(msg)
    except ValueError as e:
        print("send message error", e)

# 信息直接转发给对手
async def processTransmit(gameRoom, gameMsg):
    curUser, oppoUser = gameRoom.getCurrentUser(gameMsg.header.userId)
    if oppoUser:
        await sendMessage(oppoUser.websocket, gameMsg.toString())

    await sendMessage(curUser.websocket, gameMsg.newEchoMessage())
    return True

# 处理 room 指令消息
async def processRoom(gameRoom, gameMsg):
    return await processTransmit(gameRoom, gameMsg)

# 处理 user 指令消息
async def processUser(gameRoom, gameMsg):
    curUser, oppoUser = gameRoom.getCurrentUser(gameMsg.header.userId)
    # 更新自己的信息
    curUser.updateUser(gameMsg.data["status"], gameMsg.data["userAttribute"])
    await processTransmit(gameRoom, gameMsg)
    # 用户退出
    if curUser.isExitRoom():
        gameRoom.removeRoomUser(gameMsg.header.userId)
        game_user_dict[gameMsg.header.roomId] = gameRoom
        return
    # 收到一个用户信息交换信息
    if oppoUser:
        await sendMessage(oppoUser.websocket, gameMsg.toString())
        await sendMessage(curUser.websocket, gameMsg.newPlayUserMessage(oppoUser))
    # 所有用户都已经确认开始, 发送游戏开始消息
    if curUser and oppoUser and curUser.isStart() and oppoUser.isStart():
        gameRoom.status = 2
        gameRoom.startUpdateUser(curUser, oppoUser)
        game_user_dict[gameMsg.header.roomId] = gameRoom
        rs = round(random.random())
        orderNums = orderList(15)
        await sendMessage(curUser.websocket, gameMsg.getStartMessage(rs, curUser.userId, orderNums))
        await sendMessage(oppoUser.websocket, gameMsg.getStartMessage(1 - rs, oppoUser.userId, orderNums))

    return True

# 处理 move 指令
async def processMove(gameRoom, gameMsg):
    await processTransmit(gameRoom, gameMsg)
    return True

# 处理 over 指令
async def processOver(gameRoom, gameMsg):
    await processTransmit(gameRoom, gameMsg)
    gameRoom.overGame()
    game_user_dict[gameMsg.header.roomId] = gameRoom
    return True

# 处理 propose 提议指令
async def processPropose(gameRoom, gameMsg):
    await processTransmit(gameRoom, gameMsg)
    return True

# 核心指令分发处理
async def mainProcess(gameRoom, gameMsg):
    # 处理 room 消息
    if gameMsg.header.oper == "room":
        return await processRoom(gameRoom, gameMsg)
    # 处理用户基本信息
    if gameMsg.header.oper == "user":
        return await processUser(gameRoom, gameMsg)
    # 处理走棋信息
    if gameMsg.header.oper == "move":  
        return await processMove(gameRoom, gameMsg)
    # 处理结束信息
    if gameMsg.header.oper == "over":
        return await processOver(gameRoom, gameMsg)
    # 处理提议信息
    if gameMsg.header.oper == "propose":
        return await processPropose(gameRoom, gameMsg)

    return await processTransmit(gameRoom, gameMsg)

async def processMessage(gameMsg, websocket):
    global game_user_dict
    gameRoom = game_user_dict.get(gameMsg.header.roomId)
    # 重置对战类型
    gameMsg.resetGameBattleType(gameRoom.gameType, gameRoom.battleCate)
    isFull = gameRoom.isRoomFull(gameMsg.header.userId)
    if gameRoom.isStartGame() and isFull:
        await sendMessage(websocket, gameMsg.getRoomFull())
        return
    gameRoom = gameRoom.updateUserSocket(gameMsg.header.userId, websocket)
    game_user_dict[gameMsg.header.roomId] = gameRoom
    await mainProcess(gameRoom, gameMsg)

# 不存在房间，需要创建并插入当前用户
async def processCreateRoom(gameMsg, websocket):
    if gameMsg.header.oper != "room":
        await sendMessage(websocket, gameMsg.getRoomNonExist())
        return

    # 创建房间
    gameRoom = GameRoom(gameMsg.header)
    game_user_dict[gameMsg.header.roomId] = gameRoom
    
    await sendMessage(websocket, gameMsg.newEchoMessage())


# 接收客户端消息并处理，这里只是简单把客户端发来的返回
async def recvUserMsg(websocket):
    while True:
        recvText = await websocket.recv()
        recvJson = json.loads(recvText)
        gameMsg = GameMessage(recvJson["header"], recvJson["data"])
        if gameMsg.header.roomId in game_user_dict: # 存在房间
            await processMessage(gameMsg, websocket)
        else : # 不存在房间
           await processCreateRoom(gameMsg, websocket)

def clearRoomUser():
    for roomId in game_user_dict.keys():
        room = game_user_dict[roomId]
        for i in range(len(room.userList)):
            currentTime = time.time()
            # 每次只删除一个
            if currentTime - room.userList[i].updatetime > 180:
                print("delete user ", room, room.userList[i])
                del room.userList[i]
                game_user_dict[roomId] = room
                break
        # 删除房间
        if len(room.userList) == 0:
            print("delete room ", room)
            del game_user_dict[roomId]
            break
    
def closeCurrentWebsocket(websocket):
    isMatch = False
    for roomId in game_user_dict.keys():
        room = game_user_dict[roomId]
        for i in range(len(room.userList)):
            if room.userList[i].websocket == websocket:
                print("websocket close ", websocket)
                room.userList[i].websocket = None
                game_user_dict[roomId] = room
                isMatch = True
                break
        if isMatch :
            break

# 服务器端主逻辑
async def run(websocket, path):
    while True:
        try:
            print("recv game message ...", path)    # 链接断开
            await recvUserMsg(websocket)
        except websockets.ConnectionClosed:
            print("ConnectionClosed...", path)    # 链接断开
            closeCurrentWebsocket(websocket)
            clearRoomUser()
            break
        except websockets.InvalidState:
            print("InvalidState...")    # 无效状态
            break
        except Exception as e:
            print("Exception:", e)


if __name__ == '__main__':
    print("127.0.0.1:8881 websocket...")
    asyncio.get_event_loop().run_until_complete(websockets.serve(run, "127.0.0.1", 8881))
    asyncio.get_event_loop().run_forever()
