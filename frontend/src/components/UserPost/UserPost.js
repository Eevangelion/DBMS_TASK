import React from "react";
import { useSelector } from "react-redux";
import { Link,useLocation } from "react-router-dom";
import { useCheckIfUserSubscribedToQuery } from "../../services/service";
import "./UserPost.css";

const UserPost = (props) => {
    const location = useLocation();
    const isActive = useSelector(state => state.pagesReducer.searchPageIsActive);
    const userID = localStorage.getItem("userID");
    const userInfo = props.user;
    const {
        data: subscribed,
        isLoading
    } = useCheckIfUserSubscribedToQuery(userInfo.id);
    if (isLoading) {
        return <></>;
    }
    return (
        <div className="user-post" style={isActive ? {} : {backgroundColor: "#767676", border: "0.1vh solid #555"}}>
            <div className="user-info">
                <Link className="user-name" to={`/user/${userInfo.name}`}>{userInfo.name}</Link>
                <div className="user-subs-count">{userInfo.subscribers_count} подписчиков</div>
                <div className="user-posts-count"> {userInfo.posts_count} шуток опубликовано</div>
            </div>
                
            <Link   to={subscribed ? `/unsubscribe/${userInfo.id}` : `/subscribe/${userInfo.id}`} 
                    className="sub-button"
                    style={isActive ? {} : {backgroundColor: "#043653", color: "#aaa"}}
                    state={{ backgroundLocation: location}}
                    onClick={(event) => {if (!isActive || Number(userID) === userInfo.id) event.preventDefault()}}
            >{subscribed ? 'Отписаться' : 'Подписаться'}</Link>
        </div>
    )
}

export default UserPost;