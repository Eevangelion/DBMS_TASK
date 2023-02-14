import React, {useState} from "react";
import { useSelector } from "react-redux";
import Pagination from '@mui/material/Pagination';
import { useParams } from "react-router-dom";
import styles from "../styles/UserPage.module.css";
import JokePost from "../components/JokePost/JokePost";
import Profile from "../components/Profile/Profile";
import TopPanel from "../components/TopPanel/TopPanel";
import JokeSorter from "../components/JokeSorter/JokeSorter";
import {useGetJokesByAuthorNameQuery} from "../services/Joke";


const paginateStyle = {
    textDecoration : "none",
    color: "white",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
    fontSize: "1.4vh",
    marginLeft: "2.5vw",
    marginTop: "1vw",
}

const UserPage = (props) => {

    const [pageState, setPage] = useState(1);
    const activeButton = useSelector(state => state.buttonsReducer.sort);

    const {username} = useParams();

    const {
        data: response,
        isLoading: loadingJokes,
    } = useGetJokesByAuthorNameQuery({name: username, page: pageState, sortBy: activeButton});

    if (loadingJokes) {
        return <div>Загрузка...</div>;
    }
    const {jokes, amount} = response; 
    if (!jokes) {
        return <div className={styles.mainPage}>
                    <TopPanel />
                    <div className={styles.userInfo}>
                        <div className={styles.feed}>
                            <JokeSorter />
                            <div className={styles.txt}>Пользователь пока ничего не опубликовал</div>
                        </div>
                        <Profile username={username} />
                    </div>
                </div>;
    }
    console.log(jokes);
    const posts = jokes.map((joke) =>
    {
        return <JokePost joke={joke}/>
    });
    return (
        <div className={styles.mainPage}>
            <TopPanel />
            <div className={styles.userInfo}>
                <div className={styles.feed}>
                    <JokeSorter />
                    <div className={styles.txt}>Всего опубликовано: {amount}</div> <br/>
                    <ul className="joke-post-list">
                        {posts}
                    </ul>
                    <Pagination count={Math.ceil(amount/5)} onChange={(e, value) => setPage(value)} style={paginateStyle} shape="rounded"/>
                </div>
                <Profile username={username} />
            </div>
        </div>
    );
}

export default UserPage;