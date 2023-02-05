import React from "react";
import "../styles/UserPage.css";
import JokePost from "../components/JokePost/JokePost";
import Profile from "../components/Profile/Profile";
import TopPanel from "../components/TopPanel/TopPanel";
import JokeSorter from "../components/JokeSorter/JokeSorter";
import { useGetJokesByAuthorNameQuery } from "../services/User";

const UserPage = () => {
    const username="nikita";
    const {
        data: response,
        isLoading: loadingJokes,
    } = useGetJokesByAuthorNameQuery(username);
    const ProfileFrame = Profile(username);

    const loadingFrame = <div>Загрузка...</div>;
    const noJokesFrame = <div className="main-page">
                            <TopPanel />
                            <div className="user-info">
                                <div className="feed">
                                    <JokeSorter />
                                    <div>Пользователь пока ничего не опубликовал</div>
                                </div>
                                {Profile(username)}
                            </div>
                        </div>;
    if (loadingJokes) {
        return loadingFrame;
    }
    const {jokes, amount} = response; 
    if (!jokes) {
        return noJokesFrame;
    }
    console.log(jokes);
    const posts = jokes.map((joke) =>
    {
        // tags = getTags(joke.id);
        return <JokePost joke={joke} tags={[]}/>
    });
    return (
        <div className="main-page">
            <TopPanel />
            <div className="user-info">
                <div className="feed">
                    <JokeSorter />
                    <div>Всего опубликовано: {amount}</div> <br/>
                    <ul className="joke-post-list">
                        {posts}
                    </ul>
                </div>
                {ProfileFrame}

            </div>
        </div>
    );
}

export default UserPage;