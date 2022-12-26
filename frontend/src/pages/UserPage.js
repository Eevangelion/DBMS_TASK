import React from "react";
import "../styles/UserPage.css";
import JokePost from "../components/JokePost/JokePost";
import Profile from "../components/Profile/Profile";
import TopPanel from "../components/TopPanel/TopPanel";
import JokeSorter from "../components/JokeSorter/JokeSorter";
import { useGetJokesByAuthorNameQuery } from "../services/User";

const UserPage = () => {
    const {jokes, error} = useGetJokesByAuthorNameQuery('nikita');

    if (!jokes || error) {
        if (error && 'status' in error) {
            const errorMessage = 'error' in error ? error.error : JSON.stringify(error.data);
            return (
                <div>Error:{errorMessage}</div>
            );
        } else {
            return <div>{error?.message}</div>;
        }
    }


    const posts = jokes.map((joke) => <JokePost joke={joke} tags={[]}/>);

    return (
        <div className="main-page">
            <TopPanel />
            <div className="user-info">
            <div className="feed">
                <JokeSorter />
                <ul className="joke-post-list">
                    <div>{posts}</div>
                </ul>
            </div>
            <Profile username={'nikita'}/>

        </div>
        </div>
    );
}

export default UserPage;