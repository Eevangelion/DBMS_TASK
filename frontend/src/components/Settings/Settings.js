import { useNavigate, useLocation, Link } from "react-router-dom"
import { useState } from "react";
import { useDispatch } from "react-redux";
import { useChangePasswordMutation, useChangeUserNameMutation } from "../../services/service";
import { selectPage } from '../../store/reducers/page';
import styles from './Settings.module.css';

const linkStyle = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    textAlign: "center",
    width: "380px",
    height: "30px",
    marginTop: "2vh",
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
                        style={linkStyle}
                >
                        <strong>Редактировать список тэгов</strong>
                </Link>
                <Link   to={`/reportslist/`} 
                        style={linkStyle}
                >
                        <strong>Список жалоб</strong>
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