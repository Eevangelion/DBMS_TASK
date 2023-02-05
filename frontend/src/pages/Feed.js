import React from "react";
import "../styles/Feed.css";
import JokePost from "../components/JokePost/JokePost";
import JokeSorter from "../components/JokeSorter/JokeSorter";
import TopPanel from "../components/TopPanel/TopPanel";
import { useGetJokesByAuthorNameQuery } from "../services/User";
// import { useGetTagsByJokeIDLazyQuery } from "../services/Joke";

const Feed = () => {


    // let [getTags, {tags}] = useGetTagsByJokeIDLazyQuery();

    const username="nikita";
    const {
        data: response,
        isLoading: loadingJokes,
    } = useGetJokesByAuthorNameQuery(username);

    const loadingFrame = <div>Загрузка...</div>;

    const noJokesFrame = <div className="main-page">
                            <TopPanel />
                            <div className="feed">
                                <JokeSorter />
                                <div>Пользователь пока ничего не опубликовал</div>
                            </div>
                        </div>;

    if (loadingJokes) {
        return loadingFrame;
    }
    const {jokes, amount} = response; 
    if (!jokes) {
        return noJokesFrame;
    }

    const posts = jokes.map((joke) =>
    {
        // tags = getTags(joke.id);
        return <JokePost joke={joke} tags={[]}/>
    });

    return (
        <div className="main-page">
            <TopPanel />
            <div className="feed">
                <JokeSorter />
                <div>Всего опубликовано: {amount}</div> <br/>
                <ul className="joke-post-list">
                    {posts}
                </ul>
            </div>
        </div>
    );
}

export default Feed;