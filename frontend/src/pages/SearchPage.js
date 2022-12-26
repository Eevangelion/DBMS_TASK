import React from "react";
import {useParams} from "react-router-dom";
import "../styles/SearchPage.css";
import JokePost from "../components/JokePost/JokePost";
import TopPanel from "../components/TopPanel/TopPanel";
// import { useGetTagsByJokeIDLazyQuery } from "../services/Joke";
import { useGetJokesQuery } from "../services/Search";

const SearchPage = () => {

    const { queryArg, typeArg } = useParams();

    // let [getTags, {tags}] = useGetTagsByJokeIDLazyQuery();
    const {jokes, error} = useGetJokesQuery(queryArg, typeArg);

    if (error) {
        if ('status' in error) {
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
                <ul className="joke-post-list">
                    {posts}
                </ul>
            </div>
        </div>
    );
}

export default SearchPage;