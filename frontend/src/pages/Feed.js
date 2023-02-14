import React, {useState} from "react";
import { useSelector } from "react-redux";
import Pagination from '@mui/material/Pagination';
import styles from "../styles/Feed.module.css";
import JokePost from "../components/JokePost/JokePost";
import JokeSorter from "../components/JokeSorter/JokeSorter";
import TopPanel from "../components/TopPanel/TopPanel";
import PageSelector from "../components/PageSelector/PageSelector";
import { useGetJokesQuery } from "../services/Joke";


const paginateStyle = {
    textDecoration : "none",
    color: "white",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
    fontSize: "1.4vh",
    marginLeft: "2.5vw",
    marginTop: "1vw",
}

const Feed = (props) => {

    const [pageState, setPage] = useState(1);
    const activeButton = useSelector(state => state.buttonsReducer.sort);

    const {
        data: response,
        isLoading: loadingJokes,
    } = useGetJokesQuery({page: pageState, sortBy: activeButton});


    if (loadingJokes) {
        return <div>Загрузка...</div>;
    }
    const {jokes, amount} = response; 
    if (!jokes) {
        return <div className={styles.mainPage}>
                    <TopPanel />
                    <div className={styles.info}>
                        <div className={styles.feed}>
                            <JokeSorter />
                            <div className={styles.txt}>Никто пока ничего не публиковал</div>
                        </div>
                        <PageSelector pageState={true} />
                    </div>
                </div>;
    }

    const posts = jokes.map((joke) =>
    {
        return <JokePost joke={joke}/>
    });

    return (
        <div className={styles.mainPage}>
            <TopPanel />
            <div className={styles.info}>
                <div className={styles.feed}>
                    <JokeSorter />
                    <div className={styles.txt}>Всего опубликовано: {amount}</div> <br/>
                    <ul className={styles.jokePostList}>
                        {posts}
                    </ul>
                    <Pagination count={Math.ceil(amount/5)} onChange={(e, value) => setPage(value)} style={paginateStyle} shape="rounded"/>
                </div>
                <PageSelector pageState={true} />
            </div>
        </div>
    );
}

export default Feed;