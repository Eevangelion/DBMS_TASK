import React from "react";
import {useState, useEffect} from 'react';
import { Link, useLocation } from "react-router-dom";
import {Typography, Popover, Button} from '@mui/material';
import "./JokePost.css";
import rateImage from "../../styles/img/logo.png";
import { useGetUserByIDQuery, useGetTagsByJokeIDQuery,useAddJokeToFavoritesMutation,useRemoveJokeFromFavoritesMutation, useGetFavoritesByIDQuery } from "../../services/Joke";


const linkStyle = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    textAlign: "center",
    width: "15vw",
    height: "3vh",
    textDecoration : "none",
    border: "0.1vh solid #ccc",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
    fontSize: "1.4vh",
}

const JokePost = (props) => {
    const location = useLocation();

    const userID = localStorage.getItem("userID");
    const [anchorEl, setAnchorEl] = React.useState(null);

    const {
        data: user,
        isLoading: loadingUser,
    } = useGetUserByIDQuery(props.joke.author_id);
    const {
        data: tags,
        isLoading: loadingTags,
    } = useGetTagsByJokeIDQuery(props.joke.id);
    const [addJokeToFavorites] = useAddJokeToFavoritesMutation();
    const [removeJokeFromFavorites] = useRemoveJokeFromFavoritesMutation();

    const {
        data: favorites,
        isLoading: loadingFavorites,
    } = useGetFavoritesByIDQuery(userID);
    let addedToFavorite = false;
    if (!loadingFavorites && favorites.jokes != null) {
        for (let i = 0; i < favorites.jokes.length; i++) {
            if (favorites.jokes[i].id === props.joke.id) {
                addedToFavorite = true;
                break;
            }
        }
    }

    const handleClick = (event) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };
    const open = Boolean(anchorEl);

    if (loadingUser || loadingTags || loadingFavorites) {
        return <li>Загрузка...</li>;
    }

    const rating = props.joke.rating,
          header = props.joke.header,
          description = props.joke.description,
          author = user.name;
    let createdBy = Math.round((Date.now() - Date.parse(props.joke.creation_date))/1000);
    
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
        headerTagsFrame = (<div className="header-panel">
                                <div className="header">
                                    {header}
                                </div>
                                <div className="author">
                                    Posted by
                                    <a className="author-profile" href={'/user/' + author}>
                                        {author}
                                    </a>
                                </div>
                                <div className="create-date">
                                    {createdBy + ' ' + dmsTime} назад
                                </div>
                            </div>);
    } else {
        headerTagsFrame = (<>
                                <div className="header-panel">
                                    <div className="header">
                                        {header}
                                    </div>
                                    <div className="author">
                                        Posted by
                                        <a className="author-profile" href={'/user/' + author}>
                                            {author}
                                        </a>
                                    </div>
                                    <div className="create-date">
                                        {createdBy + ' ' + dmsTime} назад
                                    </div>
                                </div>
                                <div className="tags">
                                    {tags.map(tag => {
                                        return <div className="tag-item">{tag.name}</div>
                                    })}
                                </div>
                            </>);
    }

    return (
        <li className="joke-post">
            <div className="rating-field">
                <div className="rating">{rating}</div>
                <Link
                        className="add-to-favorite"
                        onClick={() => {
                        if (!addedToFavorite) {
                            addJokeToFavorites(props.joke.id);
                        } else {
                            removeJokeFromFavorites(props.joke.id);
                        }; addedToFavorite = !addedToFavorite;}}
                >
                    <img className="rate-image" src={rateImage} alt="?"/>
                </Link>
            </div>

            <div className="info">
                {headerTagsFrame}
                <div className="description">
                    {description.split('\n').map(str => <div>{str}<br/></div>)}
                </div>
                <div className="button-popover">
                    <button 
                        variant="contained" 
                        onClick={handleClick}
                        style={{}}>
                        ...
                    </button>
                    <Popover
                        open={open}
                        anchorEl={anchorEl}
                        onClose={handleClose}
                        anchorOrigin={{
                            vertical: 'top',
                            horizontal: 'right',
                        }}  
                        >
                        <Link style={linkStyle} state={{ backgroundLocation: location }}>Удалить шутку</Link>
                        <Link style={linkStyle} state={{ backgroundLocation: location }}>Отправить жалобу</Link>
                    </Popover>
                </div>
            </div>
        </li>
    )
}

export default JokePost;