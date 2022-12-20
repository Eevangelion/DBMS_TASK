import React, { Component } from "react";
import "../styles/Main.css";
import logo from "../styles/img/logo_test.png";
import JokePost from "./JokePost";
import TopPanelButtons from "../components/TopPanelButtons";

class Main extends Component{

    render() {
        return (<div className="main-page">
            <div className="top-panel">
                <a className="main-page-redirect" href="/">
                    <img className="main-page-redirect-image" src={logo} alt=":("/>
                </a>
                
                <div className="search-panel">
                    <form action="/search/" autoComplete="off" className="form-search" method="get" role="search">
                        <input type="search" className="search" placeholder="Поиск" />
                    </form>
                </div>
                
                {TopPanelButtons(false)}
                
            </div>
            <div className="feed">
                <JokePost/>
                <JokePost/>
            </div>
        </div>);
    }
}

export default Main;