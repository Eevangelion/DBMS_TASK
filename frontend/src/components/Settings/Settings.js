import { useNavigate, Link } from "react-router-dom"
import { useState } from "react";
import { useDispatch } from "react-redux";
import { useChangePasswordMutation, useChangeUserNameMutation } from "../../services/service";
import { selectPage } from '../../store/reducers/page';
import styles from './Settings.module.css';


const Settings = () => {
    const dispatch = useDispatch();
    dispatch(selectPage({page: 'userPage', state: false}));
    const navigate = useNavigate();
    const userRole = localStorage.getItem('userRole');
    const [usernameText, setUsernameText] = useState('');
    const [passwordText, setPasswordText] = useState('');

    const [changeName] = useChangeUserNameMutation();
    const [changePassword] = useChangePasswordMutation();

    const handleChangeUsername = (name) => {
        changeName(name);
        navigate(`/user/${name}`);
        dispatch(selectPage({page: 'userPage', state: true}));
    };
    const handleChangePassword = (password) => {
        changePassword(password);
    };
    const signOut = () => {
        localStorage.clear();
        navigate("/login");
        window.scrollTo(0,0);
    }

    if (userRole === "admin") {
        return (
        <div className={styles.modalWindow}>
            <div className={styles.modalHeader}>
                Настройки
            </div>
            <div className={styles.modalBody}>
                <div className={styles.changeUsernameForm}>
                    <text>Смена имени</text>
                    <div className={styles.changeUsername}>
                        <textarea   className={styles.newUsername} 
                                    placeholder="Введите новое имя" 
                                    onChange={e=>setUsernameText(e.target.value)} 
                                    value={usernameText} >            
                        </textarea>
                        <button 
                            className={styles.submitButton}
                            onClick={() => handleChangeUsername(usernameText)}
                        >
                            Подтвердить
                        </button>
                    </div>
                </div>
                <div className={styles.changePasswordForm}>
                    <text>Смена пароля</text> 
                    <div className={styles.changePassword}>
                        <textarea   className={styles.newPassword} 
                                    placeholder="Введите новый пароль" 
                                    onChange={e=>setPasswordText(e.target.value)} 
                                    value={passwordText} >
                        </textarea>
                        <button 
                            className={styles.submitButton}
                            onClick={() => handleChangePassword(passwordText)}
                        >
                            Подтвердить
                        </button>
                    </div>
                </div>
                <Link   to={`/tagredactor/`} 
                        className={styles.tagRedactorButton}
                >
                        <strong>Редактировать список тэгов</strong>
                </Link>
                <Link   to={`/reportslist/`} 
                        className={styles.reportListButton}
                >
                        <strong>Список жалоб</strong>
                </Link>
                <button className={styles.signOutButton} onClick={signOut}>
                    Выйти из аккаунта
                </button>
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
                    <div className={styles.changeUsernameForm}>
                        <text>Смена имени</text>
                        <div className={styles.changeUsername}>
                            <textarea   className={styles.newUsername} 
                                        placeholder="Введите новое имя" 
                                        onChange={e=>setUsernameText(e.target.value)} 
                                        value={usernameText} >            
                            </textarea>
                            <button 
                                className={styles.submitButton}
                                onClick={() => handleChangeUsername(usernameText)}
                            >
                                Подтвердить
                            </button>
                        </div>
                    </div>
                    <div className={styles.changePasswordForm}>
                        <text>Смена пароля</text> 
                        <div className={styles.changePassword}>
                            <textarea   className={styles.newPassword} 
                                        placeholder="Введите новый пароль" 
                                        onChange={e=>setPasswordText(e.target.value)} 
                                        value={passwordText} >
                            </textarea>
                            <button 
                                className={styles.submitButton}
                                onClick={() => handleChangePassword(passwordText)}
                            >
                                Подтвердить
                            </button>
                        </div>
                    </div>
                    <button className={styles.signOutButton} onClick={signOut}>
                        Выйти из аккаунта
                    </button>
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