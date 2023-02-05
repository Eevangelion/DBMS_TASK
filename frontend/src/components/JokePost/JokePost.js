import React from "react";
import "./JokePost.css";
import rateImage from "../../styles/img/logo.png";
import {useGetUserByIDQuery} from "../../services/User";


const JokePost = ({joke, tags}) => {

    const userID = localStorage.getItem("userID");
    const {
        data: user,
        isLoading: loadingUser,
    } = useGetUserByIDQuery(userID);
    
    const loadingFrame = <li>Загрузка...</li>;

    if (loadingUser) {
        return loadingFrame;
    }

    const authorName = user.name;

    const rating = joke.rating,
          header = joke.header,
          description = joke.description,
          author = authorName,
          createdBy = joke.creation_date;

    return (
        <li className="joke-post">
            <div className="rating-field">
                <div className="rating">{rating}</div>
                <a className="add-to-favorite" href="/">
                    <img className="rate-image" src={rateImage} alt="?"/>
                </a>
            </div>

            <div className="info">
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
                        {createdBy}
                    </div>
                </div>
                <div className="tags">
                    {tags.map(tag => {
                        return <div className="tag-item">{tag}</div>
                    })}
                </div>
                <div className="description">
                    {description.split('\n').map(str => <div>{str}<br/></div>)}
                </div>
            </div>
        </li>
    )
}

export default JokePost;