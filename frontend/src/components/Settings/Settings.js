import { useNavigate, useLocation, Link } from "react-router-dom"
import { useDispatch } from "react-redux";
import { useGetUserByIDQuery } from "../../services/Joke";
import { selectPage } from '../../store/reducers/page';
import styles from './Settings.module.css';

const linkStyle = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    textAlign: "center",
    width: "380px",
    height: "30px",
    borderRadius: "45px",
    backgroundColor: "#00d",
    textDecoration : "none",
    borderColor: "transparent",
    color: "white",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
}


const Settings = () => {
    const dispatch = useDispatch();
    dispatch(selectPage({page: 'userPage', state: false}));
    const location = useLocation();
    const navigate = useNavigate();
    const userID = localStorage.getItem('userID');

    const {
        data: user,
        isLoading: loadingUser, 
    }= useGetUserByIDQuery(userID);

    if (loadingUser) {
        return <div className="modal-window">Загрузка...</div>;
    }

    if (user.role === "admin") {
        return (
        <div className={styles.modalWindow}>
            <div className={styles.modalHeader}>
                Настройки
            </div>
            <div className={styles.modalBody}>
                <Link   to={`/tagredactor`} 
                        style={linkStyle}
                        state={{ backgroundLocation: location }}
                >
                        <strong>Редактировать список тэгов</strong>
                </Link>
            </div>
            <div className={styles.modalFooter}>
                <button className={styles.backButton} onClick={() => {navigate(-1);dispatch(selectPage({page: 'userPage', state: true}));}}>
                    Назад
                </button>
            </div>
        </div>);
    } else {
        return (
            <div className={styles.modalWindow}>
                <div className={styles.modalHeader}>
                    Настройки
                </div>
                <div className={styles.modalBody}>

                </div>
                <div className={styles.modalFooter}>
                    <button className={styles.backButton} onClick={() => {navigate(-1);dispatch(selectPage({page: 'userPage', state: true}));}}>
                        Назад
                    </button>
                </div>
            </div>
        );
    }
}

export default Settings;