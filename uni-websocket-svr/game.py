import json
import time
class GameUser:
    def __init__(self, userId, status, ua, websocket):
        self.userId = userId
        # 0 未知 1 已进入 2 开始  3 结束游戏 4 离开房间
        self.status = status
        self.userAttribute = ua
        self.websocket = websocket
        self.updatetime = time.time()
    
    def updateUser(self, status, ua):
        self.status = status
        self.userAttribute = ua
    
    def isExitRoom(self):
        return self.status == 4
    def isStart(self):
        return self.status == 2

class GameRoom:
    def __init__(self, header):
        self.roomId = header.roomId
        self.gameType = header.gameType
        self.battleCate = header.battleCate
        self.status = 0    # 0: 初始化完成，1: 确认中/匹配中  2: 游戏中  3 游戏结束
        self.userList = []
    def __repr__(self):
        msg = "roomId=" + self.roomId + " battleCate=" + str(self.battleCate)
        for i in range(len(self.userList)):
            msg += " " + self.userList[i].userId + " " + str(self.userList[i].status)
        return msg
    def endGame(self):
        self.status = 3
        for i in range(len(self.userList)):
            self.userList[i].status = 3
    def startGame(self):
        self.status = 2
    def isStartGame(self):
        return self.status == 2
    def friendGame(self):
        self.battleCate = 1
    def netGame(self):
        self.battleCate = 2
    
    def getCurrentUser(self, userId):
        curUser, oppoUser = None, None
        for user in self.userList:
            if user.userId != userId:
                if oppoUser == None:
                    oppoUser = user
                elif user.isStart():
                    oppoUser = user
            else :
                curUser = user
        return curUser, oppoUser

    def removeRoomUser(self, userId):
        for i in range(len(self.userList)):
            if self.userList[i].userId == userId:
                print("delete current room user ", userId)
                del self.userList[i]
                break
        return self
    
    def startUpdateUser(self, usera, userb):
        self.userList[0] = usera
        self.userList[1] = userb
        self.userList = self.userList[-2:]

    def isRoomFull(self, userId):
        for i in range(len(self.userList)):
            if self.userList[i].userId == userId:
                return False
        return len(self.userList) >= 2

    def updateUserSocket(self, userId, websocket):
        hasUser = False
        for i in range(len(self.userList)):
            if self.userList[i].userId == userId:
                self.userList[i].websocket = websocket
                self.userList[i].updatetime = time.time()
                hasUser = True
        
        if not hasUser :
            user = GameUser(userId, 0, None, websocket)
            self.userList.append(user)

        return self
    def overGame(self):
        for i in range(len(self.userList)):
            self.userList[i].status = 1
        self.endGame()

    def isAllUserStart(self):
        if len(self.userList) < 2:
            return False
        return self.userList[0].isStart() and self.userList[1].isStart()
