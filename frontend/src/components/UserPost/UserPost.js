import React from "react";
import { useSelector } from "react-redux";
import "./UserPost.css";

const UserPost = (props) => {
    const isActive = useSelector(state => state.pagesReducer.searchPageIsActive);
    const userInfo = props.user;

    return (
        <div className="user-post" style={isActive ? {} : {backgroundColor: "#767676", border: "0.1vh solid #555"}}>
            Имя {userInfo.name} <br/>
            Роль: {userInfo.role} <br/>
        </div>
    )
}

export default UserPost;