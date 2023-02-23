import React from "react";
import {useGetUserByIDQuery} from "../../services/service";
import "./TopPanelButtons.css";

const TopPanelButtons = (props) => {
    const userID = localStorage.getItem("userID");
    const {
        data: user,
        isLoading: loadingUser
    } = useGetUserByIDQuery(userID);

    if (loadingUser) {
        return (<div>Загрузка...</div>);
    }

    if (props.isAuth === false) {
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
                    <a className="profile-button" href={"/user/" + user.name}>
                        <strong>Мой профиль</strong><span style={{color: '#999'}}>{"(" + user.name + ")"}</span>
                    </a> 
                </div>);
    }
}

export default TopPanelButtons;