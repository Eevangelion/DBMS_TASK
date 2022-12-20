import React from "react";
import "../styles/JokePost.css";
import rateImage from "../styles/img/logo.png";
const JokePost = () => {

    const rating = 25,
          header = "История об одном солдате",
          description = "Чертит студент на доске окружность, а она у него ровная получается. Как будто циркулем нарисовал. Препод его спрашивает:\n- Вы где научились так окружности рисовать?\n- А я в армии два года мясорубку крутил.",
          author = "Ivan",
          createdBy = "2 часа назад";

    return (
        <div className="joke-post">
            <div className="rating-field">
                <span className="rating">{rating}</span>
                <a className="add-to-favorite" href="/">
                    <img className="rate-image" src={rateImage} alt="?"/>
                </a>
            </div>

            <div className="info">
                <div className="header-panel">
                    <div className="header">
                        <span>{header}</span>
                    </div>
                    <div className="author">
                        <span>Posted by </span>
                        <a className="author-profile" href={'/user/' + author}>
                            {author}
                        </a>
                    </div>
                    <div className="create-date">
                        <span>{createdBy}</span>
                    </div>
                </div>
                <div className="tags">
                    <div className="tag-item">Tag1</div>
                    <div className="tag-item">Tag2</div>
                </div>
                <div className="description">
                    {description.split('\n').map(str => <span>{str}<br/></span>)}
                </div>
            </div>
        </div>
    )
}

export default JokePost;