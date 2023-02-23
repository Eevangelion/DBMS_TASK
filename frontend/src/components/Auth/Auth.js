import { getAuthorizeCodeHref } from "../../store/actions/auth";
import styles from "./Auth.module.css";

const clientID = process.env.REACT_APP_CLIENT_ID;

const AuthModal = () => {
    console.log(clientID);
    const handleClick = () => {
        window.location.href = `https://github.com/login/oauth/authorize?client_id=${clientID}`;
    }

    return (
        <div className={styles.modalWindow}>
            <div className={styles.modalHeader}>
                Авторизация
            </div>
            <div className={styles.modalBody}>
                Авторизируйтесь, чтобы пользоваться сайтом
            </div>
            <div className={styles.modalFooter}>
                <button className={styles.loginButton}
                        onClick={handleClick}>
                    Авторизироваться с помощью Github
                </button>
            </div>
        </div>
    );
}

export default AuthModal;