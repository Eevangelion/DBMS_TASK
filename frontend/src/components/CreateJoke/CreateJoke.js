import {useState} from 'react';
import { useNavigate } from "react-router-dom"
import { useCreateJokeMutation } from "../../services/Joke";
import './CreateJoke.css';

const CreateJoke = () => {
    const navigate = useNavigate();

    const [createJoke] = useCreateJokeMutation();
    const [headerText, setHeaderText] = useState('');
    const [descriptionText, setDescriptionText] = useState('');
    const userID = Number(localStorage.getItem('userID'));

    const onClick = async (headerText, descriptionText) => {
        try {
            await createJoke({
                header: headerText,
                description: descriptionText,
                author_id: userID,
            });
            navigate(-1);
        } catch (error) {
            throw error;
        }
    }

    return (
        <div className="modal-window">
            <textarea className="new-header" placeholder="Заголовок вашей шутки" onChange={e=>setHeaderText(e.target.value)} value={headerText} ></textarea>
            <textarea className="new-description" placeholder="Текст вашей шутки" onChange={e=>setDescriptionText(e.target.value)} value={descriptionText} ></textarea>
            <div className="buttons">
                <button className="create-button" onClick={()=>onClick(headerText, descriptionText)}>
                    Создать
                </button>
                <button className="back-button" onClick={() => navigate(-1)}>
                    Назад
                </button>
            </div>
        </div>
    )
}

export default CreateJoke;