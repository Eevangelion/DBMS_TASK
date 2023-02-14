import React from "react";
import Pagination from '@mui/material/Pagination';
import {useParams} from "react-router-dom";
import styles from "../styles/SearchPage.module.css";
import JokePost from "../components/JokePost/JokePost";
import TopPanel from "../components/TopPanel/TopPanel";
// import { useGetTagsByJokeIDLazyQuery } from "../services/Joke";
import { useGetJokesQuery } from "../services/Search";

const paginateStyle = {
    textDecoration : "none",
    color: "white",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
    fontSize: "1.4vh",
    marginLeft: "2.5vw",
    marginTop: "1vw",
}

const SearchPage = (props) => {

    const { queryArg, typeArg } = useParams();

    // let [getTags, {tags}] = useGetTagsByJokeIDLazyQuery();
    const {jokes, error} = useGetJokesQuery(queryArg, typeArg);


    const posts = jokes.map((joke) =>
    {
        // tags = getTags(joke.id);
        return <JokePost joke={joke} tags={[]}/>
    });

    return (
        <div className={styles.mainPage}>
            <TopPanel />
            <div className={styles.feed}>
                <ul className="joke-post-list">
                    {posts}
                </ul>
            </div>
            <Pagination count={Math.ceil(5/5)} style={paginateStyle} shape="rounded"/>
        </div>
    );
}

export default SearchPage;