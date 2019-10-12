package auth

//unc SaveCurrentUserToContext(c *gin.Context) {
//user, err := getCurrentUserFromSession(c)
//if err != nil {
//return
//}
//c.Keys[config.AppConfig.ContextCurrentUserDataKey] = user
//}
//f
////// -------------- private --------------
////// getCurrentUserFromSession : 从 session 中获取用户
////func getCurrentUserFromSession(c *gin.Context) (*userModel.User, error) {
////	// 从 cookie 中获取 remember me token (如有则自动登录)
////	rememberMeToken := getRememberTokenFromCookie(c)
////	if rememberMeToken != "" {
////		if user, err := userModel.GetByRememberToken(rememberMeToken); err == nil {
////			Login(c, user)
////			return user, nil
////		}
////		delRememberToken(c)
////	}
////
////	// 从 session 中获取用户 id
////	idStr := session.GetSession(c, config.AppConfig.AuthSessionKey)
////	if idStr == "" {
////		return nil, errors.New("没有获取到 session")
////	}
////
////	id, err := strconv.Atoi(idStr)
////	if err != nil {
////		return nil, err
////	}
////
////	user, err := userModel.Get(id)
////	if err != nil {
////		return nil, err
////	}
////
////	return user, nil
////}