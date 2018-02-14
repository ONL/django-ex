# -*- coding: utf-8 -*-
# User Klasse für Login-Manager
class User:
    def __init__(self):
        self.isauthenticated = False
        self.isactive = False
        self.isanonym = True
        self.id = None
        self.name = None
        self.level = None
    
    def try_login(self, pw):
        if 'GimLummrvS!' == pw:
            self.isauthenticated = True
            self.isactive = True
            self.isanonym = False
            self.id = 1
            self.name = None
        
        
    def is_authenticated(self):
        return self.isauthenticated

    def is_active(self):
        return self.isactive
        
    def is_anonymous(self):
        return self.isanonym
        
    def get_id(self):
        return unicode(self.id)
        
    def get_name(self):
        return self.name
        
    def get_level(self):
        return self.level
    
    # Session Restore 
    def get_user(self, userid):
        self.id = int(userid)
        self.isactive = True
        self.isauthenticated = True
        self.isanonym = False
        self.name = None
