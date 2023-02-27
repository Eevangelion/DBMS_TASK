import React from "react";
import { useSelector } from "react-redux";
import { Link,useLocation } from "react-router-dom";
import "./UserPost.css";

const UserPost = (props) => {
    const location = useLocation();
    const isActive = useSelector(state => state.pagesReducer.searchPageIsActive);
    const userInfo = props.user;

    return (
        <div className="user-post" style={isActive ? {} : {backgroundColor: "#767676", border: "0.1vh solid #555"}}>
            <div className="user-info">
                <div className="user-name">{userInfo.name}</div>
                <div className="user-subs-count">{userInfo.subscribers_count} подписчиков</div>
                <div className="user-posts-count"> {userInfo.posts_count} шуток опубликовано</div>
            </div>
            <Link   to={`/subscribe/${userInfo.id}`} 
                    className="sub-button"
                    style={isActive ? {} : {backgroundColor: "#118", color: "#aaa"}}
                    state={{ backgroundLocation: location}}
                    onClick={(event) => {if (!isActive) event.preventDefault()}}
            >Подписаться</Link>
        </div>
    )
}

export default UserPost;