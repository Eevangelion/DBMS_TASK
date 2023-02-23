import React, {useState} from "react";
import { useSelector } from "react-redux";
import Pagination from '@mui/material/Pagination';
import styles from "../styles/Feed.module.css";
import JokePost from "../components/JokePost/JokePost";
import JokeSorter from "../components/Sorter/Sorter";
import TopPanel from "../components/TopPanel/TopPanel";
import PageSelector from "../components/PageSelector/PageSelector";
import { useGetJokesQuery } from "../services/service";
import { getCode } from "../store/actions/auth";
import { useGetGitQuery } from "../services/auth";


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

    const code = getCode();
    const {
        data: user,
        isLoading: loadingGavno
    } = useGetGitQuery(code);
    const [pageState, setPage] = useState(1);
    const activeButton = useSelector(state => state.buttonsReducer.sort);
    const isActive = useSelector(state => state.pagesReducer.feedIsActive);

    const {
        data: response,
        isLoading: loadingJokes,
    } = useGetJokesQuery({page: pageState, sortBy: activeButton});


    if (loadingJokes || loadingGavno) {
        return <div>Загрузка...</div>;
    }
    console.log(user); 
    const {jokes, amount} = response; 
    if (!jokes) {
        return <div className={styles.mainPage}>
                    <TopPanel />
                    <div className={styles.info} style={isActive ? {} : {backgroundColor: "#676a6c"}}>
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
            <div className={styles.info} style={isActive ? {} : {backgroundColor: "#676a6c"}}>
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