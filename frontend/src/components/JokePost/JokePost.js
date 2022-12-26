import React from "react";
import "./JokePost.css";
import rateImage from "../../styles/img/logo.png";


const JokePost = ({joke, tags}) => {

    // const rating = 25,
    //       header = "История об одном солдате",
    //       description = "Чертит студент на доске окружность, а она у него ровная получается. Как будто циркулем нарисовал. Препод его спрашивает:\n- Вы где научились так окружности рисовать?\n- А я в армии два года мясорубку крутил.",
    //       author = "Ivan",
    //       createdBy = "2 часа назад";

    const rating = joke.rating,
          header = joke.header,
          description = joke.description,
          author = joke.author,
          createdBy = joke.created_date,
          tags = tags;

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
                    <div className="tag-item">
                        Армия
                    </div>
                    <div className="tag-item">
                        Мясо
                    </div>
                </div>
                <div className="description">
                    {description.split('\n').map(str => <div>{str}<br/></div>)}
                </div>
            </div>
        </li>
    )
}

export default JokePost;