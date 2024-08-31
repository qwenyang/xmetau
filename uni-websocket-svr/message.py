import copy
import json
# 协议头部
class Header:
    def __init__(self, oper, roomId, gameId, userId, gameType, battleCate):
        self.oper = oper # 操作指令
        self.roomId = roomId # 房间ID
        self.gameId = gameId # 游戏ID，每一局/每一关的唯一ID
        self.userId = userId # 发消息的用户ID
        self.gameType = gameType # 游戏类型 翻翻棋/传统象棋/揭棋  8球/9球/斯诺克等
        self.battleCate = battleCate # 对战类型: 好友/网络/人机

class GameMessage:
    def __init__(self, header, data):
        self.header = Header(header["oper"], header["roomId"], header["gameId"], header["userId"], header["gameType"], header["battleCate"])
        self.data = data
    def __repr__(self):
        return "oper="+self.header.oper + " roomId=" + self.header.roomId + " userId=" + self.header.userId
    def resetGameBattleType(self, gameType, battleCate):
        self.header.gameType = gameType
        self.header.battleCate = battleCate
    def toString(self):
        body = {
            "header": {
                "oper": self.header.oper,
                "roomId": self.header.roomId,
                "userId": self.header.userId,
                "gameId": self.header.gameId,
                "gameType": self.header.gameType,
                "battleCate": self.header.battleCate,
            },
            "data": self.data
        }
        message = json.dumps(body)
        return message
    # 当双方用户都确认后，则发送开始游戏指令
    def getStartMessage(self, side, userId, orderNums):
        msg = {
            "header": {
                "oper": "start",
                "roomId": self.header.roomId,
                "userId": self.header.userId,
                "gameId": self.header.gameId,
                "gameType": self.header.gameType,
                "battleCate": self.header.battleCate,
            },
            "data": {
                "userId": userId,
                "side": side,
                "orderNums": orderNums,
            }
        }
        message = json.dumps(msg)
        return message
    # 房间已经满，已经存在游戏对局人数的上限了
    def getRoomFull(self):
        msg = {
            "header": {
                "oper": "room",
                "roomId": self.header.roomId,
                "userId": self.header.userId,
                "gameId": self.header.gameId,
                "gameType": self.header.gameType,
                "battleCate": self.header.battleCate,
            },
            "data": {
                "status": "full"
            }
        }
        message = json.dumps(msg)
        return message

    # 房间已销毁
    def getRoomNonExist(self):
        msg = {
            "header": {
                "oper": "room",
                "roomId": self.header.roomId,
                "userId": self.header.userId,
                "gameId": self.header.gameId,
                "gameType": self.header.gameType,
                "battleCate": self.header.battleCate,
            },
            "data": {
                "status": "nonexist",
            }
        }
        message = json.dumps(msg)
        return message
    # 心跳活跃指令
    def newEchoMessage(self):
        msg = {
            "header": {
                "oper": "echo",
                "roomId": self.header.roomId,
                "userId": self.header.userId,
                "gameId": self.header.gameId,
                "gameType": self.header.gameType,
                "battleCate": self.header.battleCate,
            },
            "data": {
                "status": "ok"
            }
        }
        message = json.dumps(msg)
        return message
    # 交换用户信息使用
    def newPlayUserMessage(self, user):
        msg = {
            "header": {
                "oper": "user",
                "roomId": self.header.roomId,
                "userId": self.header.userId,
                "gameId": self.header.gameId,
                "gameType": self.header.gameType,
                "battleCate": self.header.battleCate,
            },
            "data": {
                "status": user.status,
                "userAttribute": user.userAttribute,
            }
        }
        message = json.dumps(msg)
        return message
