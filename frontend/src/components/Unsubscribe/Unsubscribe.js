import { useNavigate } from "react-router-dom"
import { useParams } from 'react-router-dom';
import { useGetUserByIDQuery, useUnsubscribeToUserMutation } from "../../services/service";
import { selectPage } from "../../store/reducers/page";
import { useDispatch } from "react-redux";
import styles from './Unsubscribe.module.css';

const Unsubscribe = (props) => {
    const dispatch = useDispatch();
    dispatch(selectPage({page: 'userPage', state: false}));
    dispatch(selectPage({page: 'feed', state: false}));
    dispatch(selectPage({page: 'searchPage', state: false}));
    dispatch(selectPage({page: 'subscribes', state: false}));
    const navigate = useNavigate();
    const {receiverID} = useParams();
    const [subscribe] = useUnsubscribeToUserMutation();
    const {
        data: user,
        isLoading
    } = useGetUserByIDQuery(receiverID);
    const handleClick = async () => {
        await subscribe(receiverID);
        navigate(-1);
        dispatch(selectPage({page: 'userPage', state: true}));
        dispatch(selectPage({page: 'feed', state: true}));
        dispatch(selectPage({page: 'searchPage', state: true}));
        dispatch(selectPage({page: 'subscribes', state: true}));
    };

    if (isLoading) {
        return <></>;
    }
    return (
        <div className={styles.modalWindow}>
            Отписаться от пользователя {user.name}?
            <div className={styles.buttons}>
                <button className={styles.createButton} onClick={handleClick}>
                    Да
                </button>
                <button className={styles.backButton} onClick={() => {
                    navigate(-1);
                    dispatch(selectPage({page: 'userPage', state: true}));
                    dispatch(selectPage({page: 'feed', state: true}));
                    dispatch(selectPage({page: 'searchPage', state: true}));
                    dispatch(selectPage({page: 'subscribes', state: true}));
                }}>
                    Нет
                </button>
            </div>
        </div>
    )
}

export default Unsubscribe;