import React, {useEffect, useState} from "react";
import { useSelector } from "react-redux";
import Pagination from '@mui/material/Pagination';
import styles from "../styles/Feed.module.css";
import JokePost from "../components/JokePost/JokePost";
import JokeSorter from "../components/Sorter/Sorter";
import TopPanel from "../components/TopPanel/TopPanel";
import PageSelector from "../components/PageSelector/PageSelector";
import { useGetJokesQuery } from "../services/service";
import LoadingModal from "../components/LoadingModal/LoadingModal";

const Feed = (props) => {
    const [pageState, setPage] = useState(1);
    const [pageContent, setContent] = useState(<></>);
    const activeButton = useSelector(state => state.buttonsReducer.sort);
    const isActive = useSelector(state => state.pagesReducer.feedIsActive);

    const {
        data: response,
        isLoading: loadingJokes,
        error,
    } = useGetJokesQuery({page: pageState, sortBy: activeButton});
    useEffect(() => {
        if (!loadingJokes) {
            const {jokes, amount} = response; 
            if (!jokes) {
                setContent(
                    <>
                        <div className={styles.txt}>Никто пока ничего не публиковал</div>
                    </>
                );
            } else {
                const posts = jokes.map((joke) =>
                {
                    return <JokePost key={joke.id} joke={joke}/>
                });
                setContent(
                    <>
                        <div className={styles.txt}>Всего опубликовано: {amount}</div> <br/>
                        <ul className={styles.jokePostList}>
                            {posts}
                        </ul>
                    </>
                );
            }
        }
    }, [loadingJokes, response]);
    if (loadingJokes) {
        return <LoadingModal />;
    }
    if (error) {
        if (error && 'status' in error) {
            const errMsg = 'error' in error ? error.error : JSON.stringify(error.data);

            return (
                <div className="error-page">
                    <div className="error-text">
                        <div>An error has occurred:</div>
                        <div>{errMsg}</div>
                    </div>
                </div>
            );
        } else {
            return <div className="error-page">
                <div className="error-text">
                    <div>{error?.message}</div>
                </div>
            </div>;
        }
    }
    const amount = response.amount;

    return (
        <div className={styles.mainPage}>
            <TopPanel />
            <div className={styles.info} style={isActive ? {} : {backgroundColor: "#676a6c"}}>
                <div className={styles.feed}>
                    <JokeSorter />
                    {pageContent}
                    <Pagination className={styles.pagination} count={Math.ceil(amount/5)} onChange={(e, value) => setPage(value)} shape="rounded"/>
                </div>
                <PageSelector pageState={true} />
            </div>
        </div>
    );
}

export default Feed;