import React from "react";
import { useNavigate, useParams } from "react-router-dom";
import { useDispatch } from "react-redux";
import styles from "./JokeModal.module.css";
import rateImage from "../../styles/img/logo.png";
import { useGetJokeByIDQuery, useGetTagsByJokeIDQuery } from "../../services/service";
import { selectPage } from '../../store/reducers/page';

const JokeModal = (props) => {
    const dispatch = useDispatch();
    dispatch(selectPage({page: 'reportList', state: false}));
    const navigate = useNavigate();
    const {jokeID} = useParams();
    const {
        data: tags,
        isLoading: loadingTags,
    } = useGetTagsByJokeIDQuery(jokeID);
    const {
        data: joke,
        isLoading: loadingJoke,
    } = useGetJokeByIDQuery(jokeID);
    if (loadingTags || loadingJoke) {
        return <li>Загрузка...</li>;
    }

    const rating = joke.rating,
          header = joke.header,
          description = joke.description;
    let createdBy = Math.round((Date.now() - Date.parse(joke.creation_date))/1000);
    
    let dmsTime;

    if (createdBy >= 86400) {
        createdBy = Math.round(createdBy / 86400);
        switch (createdBy % 10) {
        case 1: dmsTime = "день"; break;
        case 2: case 3: case 4: dmsTime="дня"; break;
        default: dmsTime="дней";break;
        }
    } else if (createdBy >= 3600) {
        createdBy = Math.round(createdBy / 3600);
        switch (createdBy % 10) {
        case 1: dmsTime = "час"; break;
        case 2: case 3: case 4: dmsTime="часа"; break;
        default: dmsTime="часов";break;
        }
    } else if (createdBy >= 60) {
        createdBy = Math.round(createdBy / 60);
        switch (createdBy % 10) {
        case 1: dmsTime = "минута"; break;
        case 2: case 3: case 4: dmsTime="минуты"; break;
        default: dmsTime="минут";break;
        }
    } else {
        switch (createdBy % 10) {
        case 1: dmsTime = "секунда"; break;
        case 2: case 3: case 4: dmsTime="секунды"; break;
        default: dmsTime="секунд";break;
        }
    }

    let headerTagsFrame;

    if (!tags) {
        headerTagsFrame = (<div className={styles.headerPanel}>
                                <div className={styles.header}>
                                    {header}
                                </div>
                                <div className={styles.createDate}>
                                    {createdBy + ' ' + dmsTime} назад
                                </div>
                            </div>);
    } else {
        headerTagsFrame = (<>
                                <div className={styles.headerPanel}>
                                    <div className={styles.header}>
                                        {header}
                                    </div>
                                    <div className={styles.createDate}>
                                        {createdBy + ' ' + dmsTime} назад
                                    </div>
                                </div>
                                <div className={styles.tags}>
                                    {tags.map(tag => {
                                        return <div className={styles.tagItem}>{tag.name}</div>
                                    })}
                                </div>
                            </>);
    }

    return (
        <div className={styles.modalWindow} key={props.key}>
            <div className={styles.modalHeader}>
                Шутка 
                <div id={styles.mdiv} onClick={() => {
                        navigate(-1);
                        dispatch(selectPage({page: 'reportList', state: true}));
                    }
                }>
                    <div class={styles.mdiv}>
                        <div class={styles.md}></div>
                    </div>
                </div>
            </div>
            <div className={styles.modalBody}>
                <div className={styles.ratingField} >
                    <div className={styles.rating}>{rating}</div>
                        <div className={styles.addToFavorite}>
                            <img className={styles.rateImage} src={rateImage} alt="?"/>
                        </div>
                </div>

                <div className={styles.info}>
                    {headerTagsFrame}
                    <div className={styles.description}>
                        {description.split('\n').map(str => <div>{str}<br/></div>)}
                    </div>
                </div>
            </div>
        </div>
    )
}

export default JokeModal;