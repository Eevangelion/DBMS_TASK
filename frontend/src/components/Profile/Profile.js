import React from "react";
import { Link, useLocation } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { useGetUserByNameQuery } from "../../services/service";
import LoadingModal from "../LoadingModal/LoadingModal";
import "./Profile.css";

const Profile = (props) => {
    const userPageIsActive = useSelector(state => state.pagesReducer.userPageIsActive);
    const location = useLocation();
    const {
        data: user,
        isLoading: loadingUser,
    } = useGetUserByNameQuery(props.username);
    if (loadingUser) {
        return <LoadingModal />;
    }
    if (!user) {
        return <div className="profile-block">Пользователя с таким именем не существует</div>;
    }
    let userAccount = true;
    if (props.username !== localStorage.getItem("userName")) {
        userAccount = false;
    }
    
    const reports = user.reports,
          lastUnbanDate = user.unban_date,
          role = user.role;
    return (
        <div className="profile-block" style={userPageIsActive ? {} : {backgroundColor: "#767676", border: "0.1vh solid #555"}}>
            <strong>{props.username}</strong>
            <strong style={{color: "#999"}}>user/{props.username}</strong>
            {
                userAccount ?
                <div className={"settings-link"}>
                    <Link   to={`/settings`} 
                            className={userPageIsActive ? "link" : "link-disabled"}
                            state={{ backgroundLocation: location }}
                            onClick={(event) => {if (!userPageIsActive) event.preventDefault()}}
                    >
                        <strong>Настройки</strong>
                    </Link> 
                </div> : <></>
            }
            <div className="profile-info">
                Роль: {role === "admin" ? "Администратор" : "Пользователь"}<br/>
                Жалобы: {reports}<br/>
                {/* Добавлено в избранное: {addedToFavorite.length} <br/> */}
                Последняя дата разблокировки: {(Date.now() - Date.parse(lastUnbanDate))/1000}
            </div>
            {
                userAccount ?
                <div className="post-joke">
                    <Link   to={`/create_joke`} 
                            className={userPageIsActive ? "link" : "link-disabled"}
                            state={{ backgroundLocation: location }}
                            onClick={(event) => {if (!userPageIsActive) event.preventDefault()}}
                    >
                        <strong>Создать шутку</strong>
                    </Link>
                </div> : 
                <div className="subscribe">
                    <Link   to={`/subscribe/${user.id}`}
                            className={userPageIsActive ? "link" : "link-disabled"}
                            state={{ backgroundLocation: location}}
                            onClick={(event) => {if (!userPageIsActive) event.preventDefault()}}
                    >
                        <strong>Подписаться</strong>
                    </Link> 
                </div>
            }
        </div>
    );
}

export default Profile;