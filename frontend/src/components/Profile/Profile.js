import React, {useState, useEffect} from "react";
import { Link, useLocation } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { useGetUserByNameQuery, useCheckIfUserSubscribedToQuery } from "../../services/service";
import "./Profile.css";

const Profile = (props) => {
    const [userID, setUserID] = useState(0);
    const userPageIsActive = useSelector(state => state.pagesReducer.userPageIsActive);
    const location = useLocation();
    const currentUserRole = localStorage.getItem("userRole");
    const {
        data: user,
        isLoading: loadingUser,
    } = useGetUserByNameQuery(props.username);

    useEffect(()=> {
        if (user)
            setUserID(user.id);
    }, [user]);

    const {
        data: subscribed,
        isLoading: loadingSubscribed,
    } = useCheckIfUserSubscribedToQuery(userID);

    if (loadingUser || loadingSubscribed) {
        return <></>;
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
            <strong style={userPageIsActive ? {color: "#999"} : {color: "#666"}}>user/{props.username}</strong>
            {
                userAccount ?
                <div className={"settings-link"}>
                    <Link   to={currentUserRole === 'admin' ? '/develop_settings' : `/settings`} 
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
                Последняя дата разблокировки: {lastUnbanDate.split('T')[0]}
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
                    <Link   to={subscribed ? `/unsubscribe/${user.id}` : `/subscribe/${user.id}`} 
                            className={userPageIsActive ? "link" : "link-disabled"}
                            state={{ backgroundLocation: location}}
                            onClick={(event) => {if (!userPageIsActive) event.preventDefault()}}
                    >
                        <strong>{subscribed ? 'Отписаться' : 'Подписаться'}</strong>
                    </Link> 
                </div>
            }
        </div>
    );
}

export default Profile;