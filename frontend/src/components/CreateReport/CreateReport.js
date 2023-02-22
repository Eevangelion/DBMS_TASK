import {useState} from 'react';
import { useDispatch } from 'react-redux';
import { useNavigate } from "react-router-dom"
import { useParams } from 'react-router-dom';
import { useCreateReportMutation } from "../../services/Joke";
import { selectPage } from '../../store/reducers/page';
import styles from './CreateReport.module.css';

const CreateReport = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const {jokeID} = useParams();

    dispatch(selectPage({page: 'userPage', state: false}));
    dispatch(selectPage({page: 'feed', state: false}));
    dispatch(selectPage({page: 'searchPage', state: false}));
    dispatch(selectPage({page: 'subscribes', state: false}));

    const [createReport] = useCreateReportMutation();
    const [descriptionText, setDescriptionText] = useState('');

    const handleClick = async (descriptionText) => {
        try {
            await createReport({
                description: descriptionText,
                jokeID: Number(jokeID),
            });

            navigate(-1);
            dispatch(selectPage({page: 'userPage', state: true}));
            dispatch(selectPage({page: 'feed', state: true}));
            dispatch(selectPage({page: 'searchPage', state: true}));
            dispatch(selectPage({page: 'subscribes', state: true}));
        } catch (error) {
            throw error;
        }
    }

    return (
        <div className={styles.modalWindow}>
            <div className={styles.modalHeader}>
                Создание жалобы
            </div>
            <div className={styles.modalBody}>
            <textarea className={styles.newDescription} placeholder="Описание жалобы" onChange={e=>setDescriptionText(e.target.value)} value={descriptionText} ></textarea>
            </div>
            <div className={styles.modalFooter}>
                <button className={styles.createButton} onClick={()=>handleClick(descriptionText)}>
                    Отправить
                </button>
                <button className={styles.backButton} onClick={() => {
                    navigate(-1);
                    dispatch(selectPage({page: 'userPage', state: true}));
                    dispatch(selectPage({page: 'feed', state: true}));
                    dispatch(selectPage({page: 'searchPage', state: true}));
                    dispatch(selectPage({page: 'subscribes', state: true}));
                }}>
                    Назад
                </button>
            </div>
        </div>
    )
}

export default CreateReport;