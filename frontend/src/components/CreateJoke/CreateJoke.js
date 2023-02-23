import {useState} from 'react';
import { FormControl, InputLabel, Select, MenuItem } from '@mui/material';
import { useNavigate } from "react-router-dom";
import { useDispatch } from 'react-redux';
import { useCreateJokeMutation, useGetTagsQuery } from "../../services/service";
import styles from './CreateJoke.module.css';
import { selectPage } from '../../store/reducers/page';

const CreateJoke = (props) => {
    const dispatch = useDispatch();
    dispatch(selectPage({page: 'userPage', state: false}));
    const navigate = useNavigate();

    const [createJoke] = useCreateJokeMutation();
    const [headerText, setHeaderText] = useState('');
    const [descriptionText, setDescriptionText] = useState('');
    const [currentTags, setTags] = useState([]);
    const {
        data: tags,
        isLoading: loadingTags
    } = useGetTagsQuery();
    const userID = Number(localStorage.getItem('userID'));

    if (loadingTags) {
        return <div>Загрузка...</div>;
    }

    const handleClick = async (headerText, descriptionText, tags) => {
        try {
            await createJoke({
                header: headerText,
                description: descriptionText,
                author_id: userID,
                tags: tags,
            });

            navigate(-1);
            dispatch(selectPage({page: 'userPage', state: true}));
        } catch (error) {
            throw error;
        }
    }
    
    const handleChange = (tag) => {
        setTags(arr => {
            return (arr.indexOf(tag) !== -1) ? arr : [...arr, tag]});
    }
    const handleClickTag = (event) => {
        setTags(currentTags.filter(tag => tag.name !== event.target.innerText));
    }
    const tagsFrames = tags.map((tag) => {return <MenuItem value={tag}>{tag.name}</MenuItem>;});
    if (currentTags.length !== 0) {
        const tagsItems = currentTags.map((tag) => { return <div className={styles.tagItem} onMouseEnter={(event)=>{event.target.style.cursor = 'pointer';}}>{tag.name}</div>;});
        return (
            <div className={styles.modalWindow}>
                <div className={styles.modalHeader}>
                    Создание шутки
                </div>
                <div className={styles.modalBody}>
                    <textarea className={styles.newHeader} placeholder="Заголовок вашей шутки" onChange={e=>setHeaderText(e.target.value)} value={headerText} ></textarea>
                    <textarea className={styles.newDescription} placeholder="Текст вашей шутки" onChange={e=>setDescriptionText(e.target.value)} value={descriptionText} ></textarea>
                    <div className={styles.tags} onClick={handleClickTag}>
                        {tagsItems}
                    </div>
                    <div className={styles.addTags}>
                    <div style={{fontSize: "1vw", marginRight: "1vw"}}>Добавьте тэги: </div>
                        <FormControl style={{width: "10vw", height:"6vh", fontSize: "0.8vw"}}>
                            <InputLabel>Тэг</InputLabel>
                            <Select
                                onChange={(event) => {
                                        console.log(event.target);
                                        handleChange(event.target.value)
                                    }
                                }
                            >
                                {tagsFrames}
                            </Select>
                        </FormControl>
                    </div>
                </div>
                <div className={styles.modalFooter}>
                    <button className={styles.createButton} onClick={()=>handleClick(headerText, descriptionText, currentTags)}>
                        Создать
                    </button>
                    <button className={styles.backButton} onClick={() => {dispatch(selectPage({page: 'userPage', state: true})); navigate(-1);}}>
                        Назад
                    </button>
                </div>
            </div>
        )
    } else {
        return (
            <div className={styles.modalWindow}>
                <div className={styles.modalHeader}>
                    Создание шутки
                </div>
                <div className={styles.modalBody}>
                    <textarea className={styles.newHeader} placeholder="Заголовок вашей шутки" onChange={e=>setHeaderText(e.target.value)} value={headerText} ></textarea>
                    <textarea className={styles.newDescription} placeholder="Текст вашей шутки" onChange={e=>setDescriptionText(e.target.value)} value={descriptionText} ></textarea>
                    <div className={styles.addTags}>
                    <div style={{fontSize: "1vw", marginRight: "1vw"}}>Добавьте тэги: </div>
                        <FormControl style={{width: "10vw", height:"6vh", fontSize: "0.8vw"}}>
                            <InputLabel>Тэг</InputLabel>
                            <Select
                                onChange={(event) => {
                                        console.log(event.target);
                                        handleChange(event.target.value)
                                    }
                                }
                            >
                                {tagsFrames}
                            </Select>
                        </FormControl>
                    </div>
                </div>
                <div className={styles.modalFooter}>
                    <button className={styles.createButton} onClick={()=>handleClick(headerText, descriptionText, currentTags)}>
                        Создать
                    </button>
                    <button className={styles.backButton} onClick={() => {dispatch(selectPage({page: 'userPage', state: true})); navigate(-1);}}>
                        Назад
                    </button>
                </div>
            </div>
        )
    }
}

export default CreateJoke;