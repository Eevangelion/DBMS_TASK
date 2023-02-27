import React, {useState} from "react";
import { useSelector } from "react-redux";
import Pagination from '@mui/material/Pagination';
import {useParams, useSearchParams} from "react-router-dom";
import styles from "../styles/SearchPage.module.css";
import JokePost from "../components/JokePost/JokePost";
import Sorter from "../components/Sorter/Sorter";
import TopPanel from "../components/TopPanel/TopPanel";
import { useGetSearchResultQuery } from "../services/service";
import UserPost from "../components/UserPost/UserPost";
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

    const [pageState, setPage] = useState(1);
    const activeButton = useSelector(state => state.buttonsReducer.sort);
    const isActive = useSelector(state => state.pagesReducer.searchPageIsActive);

    const [searchParams] = useSearchParams();
    const { type: typeArg } = useParams();
    const queryArg = searchParams.get('query');


    const {
        data: response,
        isLoading: loadingSearch,
    } = useGetSearchResultQuery({q: queryArg, t: typeArg, page: pageState, sortBy: activeButton});

    if (loadingSearch) {
        return <div className={styles.mainPage}>Загрузка...</div>;
    }
    if (typeArg === 'keyword' || typeArg === 'tag') {
        const {jokes, amount} = response; 

        if (!jokes) {
            return <div className={styles.mainPage}>
                        <TopPanel />
                        <div className={styles.info} style={isActive ? {} : {backgroundColor: "#676a6c"}}>
                            <div className={styles.feed}>
                                <Sorter />
                                <div className={styles.txt}>По данному запросу ничего не найдено</div>
                                <Pagination count={Math.ceil(amount/5)} onChange={(e, value) => setPage(value)} style={paginateStyle} shape="rounded"/>
                            </div>
                        </div>
                    </div>;
        }

        const posts = jokes.map((joke) =>
        {
            return <JokePost key={joke.id} joke={joke}/>
        });
        return (
            <div className={styles.mainPage}>
                <TopPanel />
                <div className={styles.info} style={isActive ? {} : {backgroundColor: "#676a6c"}}>
                    <div className={styles.feed}>
                        <Sorter />
                        <div className={styles.txt}>Результаты поиска по {typeArg === 'keyword' ? `ключевому слову ${queryArg}` : `тэгу ${queryArg}`}</div> <br/>
                        <ul className={styles.jokePostList}>
                            {posts}
                        </ul>
                        <Pagination count={Math.ceil(amount/5)} onChange={(e, value) => setPage(value)} style={paginateStyle} shape="rounded"/>
                    </div>
                </div>
            </div>
        );
    } else {
        const people = response; 
        const amount = people.length;

        if (!people) {
            return <div className={styles.mainPage}>
                        <TopPanel />
                        <div className={styles.info} style={isActive ? {} : {backgroundColor: "#676a6c"}}>
                            <div className={styles.feed}>
                                <div className={styles.txt}>По данному запросу ничего не найдено</div>
                                <Pagination count={Math.ceil(amount/5)} onChange={(e, value) => setPage(value)} style={paginateStyle} shape="rounded"/>
                            </div>
                        </div>
                    </div>;
        }
        const posts = people.map((user) =>
        {
            return <UserPost key={user.id} user={user}/>
        });
        return (
            <div className={styles.mainPage}>
                <TopPanel />
                <div className={styles.info} style={isActive ? {} : {backgroundColor: "#676a6c"}}>
                    <div className={styles.feed}>
                        <div className={styles.txt}>Результаты поиска пользователей по имени</div> <br/>
                        <ul className={styles.peoplePostList}>
                            {posts}
                        </ul>
                        <Pagination count={Math.ceil(amount/5)} onChange={(e, value) => setPage(value)} style={paginateStyle} shape="rounded"/>
                    </div>
                </div>
            </div>
        );
    }
}

export default SearchPage;