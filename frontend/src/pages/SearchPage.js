import React, { Component } from "react";
import "../styles/Main.css";
import JokePost from "../components/JokePost/JokePost";
import JokeSorter from "../components/JokeSorter/JokeSorter";
import TopPanel from "../components/TopPanel/TopPanel";

class SearchPage extends Component{

    render() {
        return (
        <div className="main-page">
            <TopPanel />
            <div className="feed">
                <JokeSorter />
                <ul className="joke-post-list">
                    <JokePost/>
                    <JokePost/>
                </ul>
            </div>
        </div>);
    }
}

export default SearchPage;