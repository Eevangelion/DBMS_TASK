import React, { Component } from "react";
import "../styles/Main.css";
import JokePost from "../components/JokePost/JokePost";
import JokeSorter from "../components/JokeSorter/JokeSorter";
import TopPanel from "../components/TopPanel/TopPanel";
import { useGetJokesByAuthorNameQuery } from "../services/User";

const Main = () => {

    const {jokes, error} = useGetJokesByAuthorNameQuery();

    if (error || jokes === undefined) {
        if (error && 'status' in error) {
            const errorMessage = 'error' in error ? error.error : JSON.stringify(error.data);
            return (
                <div>
                    <div>Error:{errorMessage}</div>
                </div>
            );
        } else {
            return <div>{error?.message}</div>;
        }
    }

    const posts = Array.from(jokes.map((joke)=><JokePost joke={joke} />));

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
        </div>
    );
}

export default Main;