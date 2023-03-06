import React, {useEffect, useState} from "react";
import { useSelector } from "react-redux";
import Pagination from '@mui/material/Pagination';
import PageSelector from "../components/PageSelector/PageSelector";
import styles from "../styles/Subscribes.module.css";
import JokePost from "../components/JokePost/JokePost";
import JokeSorter from "../components/Sorter/Sorter";
import TopPanel from "../components/TopPanel/TopPanel";
import { useGetSubscribedByIDQuery } from "../services/service";
import LoadingModal from "../components/LoadingModal/LoadingModal";


const Subscribes = (props) => {

    const [pageState, setPage] = useState(1);
    const [pageContent, setContent] = useState(<></>);
    const activeButton = useSelector(state => state.buttonsReducer.sort);
    const isActive = useSelector(state => state.pagesReducer.subscribesIsActive);
    const userID = localStorage.getItem("userID");

    const {
        data: response,
        isLoading: loadingJokes,
        error,
    } = useGetSubscribedByIDQuery({id: userID, page: pageState, sortBy: activeButton});

    useEffect(() => {
        if (!loadingJokes) {
            if (response) {
                const {jokes, amount} = response;
                if (!jokes) {
                    setContent(
                        <>
                            <div className={styles.txt}>Пользователи, на которых вы подписаны, пока ничего не опубликовали</div>
                        </>
                    );
                }
                else {
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
            } else {
                setContent(
                    <>
                        <div className={styles.txt}>Пользователи, на которых вы подписаны, пока ничего не опубликовали</div>
                    </>
                );
            }
        }   
    }, [response, loadingJokes]);
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
    const amount = response ? response.amount : 0; 

    return (
        <div className={styles.mainPage}>
            <TopPanel />
            <div className={styles.info} style={isActive ? {} : {backgroundColor: "#676a6c"}}>
                <div className={styles.feed}>
                    <JokeSorter />
                    {pageContent}
                    <Pagination className={styles.pagination} count={Math.ceil(amount/5)} onChange={(e, value) => setPage(value)} shape="rounded"/>
                </div>
                <PageSelector pageState={false} />
            </div>
        </div>
    );
}

export default Subscribes;