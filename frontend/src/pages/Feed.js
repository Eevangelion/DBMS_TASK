import React from "react";
import "../styles/Feed.css";
import JokePost from "../components/JokePost/JokePost";
import JokeSorter from "../components/JokeSorter/JokeSorter";
import TopPanel from "../components/TopPanel/TopPanel";
import { useGetJokesByAuthorNameQuery } from "../services/User";
// import { useGetTagsByJokeIDLazyQuery } from "../services/Joke";

const Feed = () => {


    // let [getTags, {tags}] = useGetTagsByJokeIDLazyQuery();

    const {jokes, error} = useGetJokesByAuthorNameQuery();

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
                <ul className="joke-post-list">
                    <div>{posts}</div>
                </ul>
            </div>
        </div>
    );
}

export default Feed;