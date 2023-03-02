import React, {useState} from "react";
import { Link } from 'react-router-dom';
import { useSelector } from "react-redux";
import styles from "./PageSelector.module.css";


const PageSelector = (props) => {
    const userPageIsActive = useSelector(state => state.pagesReducer.userPageIsActive);
    const feedIsActive = useSelector(state => state.pagesReducer.feedIsActive);
    const searchPageIsActive = useSelector(state => state.pagesReducer.searchPageIsActive);
    const subscribesIsActive = useSelector(state => state.pagesReducer.subscribesIsActive);
    const isActive = (
        userPageIsActive &&
        feedIsActive &&
        searchPageIsActive &&
        subscribesIsActive
    );
    const [pageState, setPage] = useState(props.pageState);
    
    return (
        <div className={styles.pageSelector}>
            {isActive ? (pageState ?
            <Link   
                to={`/feed/`} 
                className={styles.linkDisabled}
                onClick={ (event) => event.preventDefault() }
            >Все шутки</Link>
            :   <Link   to={`/feed/`} 
                        className={styles.link}
                        onClick={() => (setPage(false))}
            >Все шутки</Link>) :
            <Link   to={`/feed/`}
                    className={pageState ? styles.linkDisabled : styles.link}
                    style={pageState ? { backgroundColor: "#444", color: "#555"} : {backgroundColor: "#043653", color: "#666"}}
                    onClick={(event)=>event.preventDefault()}
            >Все шутки</Link>}
            {isActive ? (pageState ? 
            <Link   
                to={`/subscribes/`}
                className={styles.link}
                onClick={ () => setPage(true) }
            >Подписки</Link>
            :   <Link   to={`/subscribes/`}
                        className={styles.linkDisabled}
                        onClick={ (event) => event.preventDefault() }
            >Подписки</Link>) : 
            <Link   to={`/subscribes/`}
                    className={pageState ? styles.linkDisabled : styles.link}
                    style={pageState ? { backgroundColor: "#043653", color: "#666"} : {backgroundColor: "#444", color: "#555"}}
                    onClick={(event)=>event.preventDefault()}
            >Подписки</Link>}
        </div>
    );
}

export default PageSelector;