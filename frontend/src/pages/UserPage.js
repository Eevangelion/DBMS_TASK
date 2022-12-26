import React, { Component } from "react";
import "../styles/UserPage.css";
import JokePost from "../components/JokePost/JokePost";
import Profile from "../components/Profile/Profile";
import TopPanel from "../components/TopPanel/TopPanel";
import JokeSorter from "../components/JokeSorter/JokeSorter";

class UserPage extends Component{
    render() {
        return (<div className="main-page">
            <TopPanel />
            <div className="user-info">
                <div className="feed">
                    <JokeSorter />
                    <ul className="joke-post-list">
                        <JokePost/>
                        <JokePost/>
                    </ul>
                </div>
                <Profile />

            </div>
        </div>);
    }
}

export default UserPage;