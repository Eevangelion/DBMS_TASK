import React from "react";
import "../styles/TopPanelButtons.css";

function TopPanelButtons(isAuth) {
    if (isAuth === false) {
        return (<div className="auth">
                    <div className="login">
                        <a className="login-button" href="/login">Войти</a> 
                    </div>
                    <div className="register">
                        <a className="register-button" href="/register">Зарегистрироваться</a>
                    </div>
                </div>)
    } else {
        return (<div className="profile">
                    <a className="profile-button" href='user/Ivan'>
                        <strong>My profile</strong><span style={{color: '#999'}}>{'(Ivan)'}</span>
                    </a> 
                </div>);
    }
}

export default TopPanelButtons;