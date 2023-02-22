import React, {useState} from "react";
import { useSelector } from "react-redux";
import Pagination from '@mui/material/Pagination';
import PageSelector from "../components/PageSelector/PageSelector";
import styles from "../styles/Subscribes.module.css";
import JokePost from "../components/JokePost/JokePost";
import JokeSorter from "../components/Sorter/Sorter";
import TopPanel from "../components/TopPanel/TopPanel";
import { useGetSubscribedByIDQuery } from "../services/Joke";



const paginateStyle = {
    textDecoration : "none",
    color: "white",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
    fontSize: "1.4vh",
    marginLeft: "2.5vw",
    marginTop: "1vw",
}

const Subscribes = (props) => {

    const [pageState, setPage] = useState(1);
    const activeButton = useSelector(state => state.buttonsReducer.sort);
    const isActive = useSelector(state => state.pagesReducer.subscribesIsActive);
    const userID = localStorage.getItem("userID");

    const {
        data: response,
        isLoading: loadingJokes,
    } = useGetSubscribedByIDQuery({id: userID, page: pageState, sortBy: activeButton});

    if (loadingJokes) {
        return <div>Загрузка...</div>;
    }
    const {jokes, amount} = response; 
    if (!jokes) {
        return <div className={styles.mainPage}>
                    <TopPanel />
                    <div className={styles.info} style={isActive ? {} : {backgroundColor: "#676a6c"}}>
                        <div className={styles.feed}>
                            <JokeSorter />
                            <div className={styles.txt}>Пользователи, на которых вы подписаны, пока ничего не опубликовали</div>
                        </div>
                        <PageSelector pageState={false}/>
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
                <PageSelector pageState={false} />
            </div>
        </div>
    );
}

export default Subscribes;