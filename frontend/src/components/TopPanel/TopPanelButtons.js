import React from "react";
import "./TopPanelButtons.css";

function TopPanelButtons(isAuth) {
    const username = "nikita";
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
                    <a className="profile-button" href={"/user/" + username}>
                        <strong>Мой профиль</strong><span style={{color: '#999'}}>{"(" + username + ")"}</span>
                    </a> 
                </div>);
    }
}

export default TopPanelButtons;