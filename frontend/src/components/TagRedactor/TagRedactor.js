import {useState} from 'react';
import { useNavigate } from "react-router-dom";
import { useCreateTagMutation, useDeleteTagMutation, useGetTagsQuery } from "../../services/Joke";
import styles from './TagRedactor.module.css';

const TagRedactor = () => {
    const navigate = useNavigate();
    const userID = localStorage.getItem("userID");

    const [createTag] = useCreateTagMutation();
    const [deleteTag] = useDeleteTagMutation();
    const [addedTags, setAddedTags] = useState([]);
    const [removedTags, setRemovedTags] = useState([]);
    const [currentTags, setTags] = useState([]);
    const [newTagName, setTagName] = useState("");
    const {
        data: tags,
        isLoading: loadingTags
    } = useGetTagsQuery();
    if (loadingTags) {
        return <div className={styles.modalWindow}>Загрузка...</div>;
    }
    if (currentTags.length === 0) {
        setTags(tags);
        for (let i = 0; i < addedTags.length; i++) {
            setTags(arr => [...arr, addedTags[i]]);
        }
        for (let i = 0; i < removedTags.length; i++) {
            setTags(currentTags.filter(tag => tag.name !== removedTags[i].name));
        }
    }
    console.log(currentTags);
    const handleClickCreateButton = (event) => {
        for (let i = 0; i < currentTags.length; i++) {
            if (currentTags[i].name === newTagName) {
                return;
            }
        }
        setAddedTags(arr => [...arr, {id: -1, name: newTagName}]);
        setTags(arr => [...arr, {id: -1, name: newTagName}]);
    }

    const handleClickSaveButton = () => {
        for (let i = 0; i < addedTags.length; i++) {
            try {
                createTag({name: addedTags[i].name, id: userID});
            } catch (err) {
                console.log(err);
            }
        }
        for (let i = 0; i < removedTags.length; i++) {
            try {
                deleteTag({name: removedTags[i].name, id: userID});
            } catch (err) {
                console.log(err);
            }
        }
        setAddedTags([]);
        setRemovedTags([]);
    }

    const handleClickTag = (event) => {
        const tagName = event.target.innerText;
        for (let i = 0; i < currentTags.length; i++) {
            if (currentTags[i].name === tagName) {
                if (currentTags[i].id === -1) {
                    setAddedTags(addedTags.filter(tag => tag.name !== tagName));
                } else {
                    setRemovedTags(arr => [...arr, currentTags[i]]);
                }
                setTags(currentTags.filter(tag => tag.name !== tagName));
                break;
            }
        }
    }

    const tagsItems = currentTags.map((tag) => { return <div className={styles.tagItem} onMouseEnter={(event)=>{event.target.style.cursor = 'pointer';}}>{tag.name}</div>;});

    return (
        <div className={styles.modalWindow}>
            <div className={styles.tags} onClick={handleClickTag}>
                {tagsItems}
            </div>
            <input className={styles.newTag} placeholder="Название тэга" onChange={e=>setTagName(e.target.value)} value={newTagName} ></input>
                <div className={styles.buttons}>
                <button className={styles.createButton} onClick={handleClickCreateButton}>
                    Создать новый тэг
                </button>
                <button className={styles.saveButton} onClick={handleClickSaveButton}>
                    Сохранить
                </button>
                <button className={styles.backButton} onClick={() => navigate(-1)}>
                    Назад
                </button>
            </div>
        </div>
    )

}

export default TagRedactor;