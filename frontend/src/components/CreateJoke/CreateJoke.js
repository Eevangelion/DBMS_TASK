import { useNavigate } from "react-router-dom"
import { useCreateJokeMutation } from "../../services/Joke";
import './CreateJoke.css';

const CreateJoke = () => {
    const navigate = useNavigate();

    const [createJoke] = useCreateJokeMutation();
    let headerText, 
        descriptionText;
    const userID = localStorage.getItem('userID');

    const onClick = async (headerText, descriptionText) => {
        try {
            await createJoke({
                header: headerText,
                description: descriptionText,
                author_id: userID,
            });
        } catch (error) {
            throw error;
        }
        navigate(-1);
    }

    return (
        <div className="modal-window">
            <input type="text" className="new-header" placeholder="Заголовок вашей шутки" value={headerText}></input>
            <textarea className="new-description" placeholder="Текст вашей шутки" value={descriptionText}></textarea>
            <div className="buttons">
                <button className="create-button" onClick={onClick}>
                    Создать
                </button>
                <button className="back-button" onClick={navigate(-1)}>
                    Назад
                </button>
            </div>
        </div>
    )
}

export default CreateJoke;