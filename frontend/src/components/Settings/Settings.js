import { useNavigate } from "react-router-dom"
import { useState } from "react";
import { useDispatch } from "react-redux";
import { useChangePasswordMutation, useChangeUserNameMutation } from "../../services/service";
import { selectPage } from '../../store/reducers/page';
import styles from './Settings.module.css';
import { AuthContext } from "../../context/context";


const Settings = () => {
    const dispatch = useDispatch();
    dispatch(selectPage({page: 'userPage', state: false}));
    const navigate = useNavigate();
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
    return (
        <div className={styles.modalWindow}>
            <div className={styles.modalHeader}>
                Настройки
            </div>
            <div className={styles.modalBody}>
                <div className={styles.changeUsernameForm}>
                    <p>Смена имени</p>
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
                    <p>Смена пароля</p> 
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
                <AuthContext.Consumer>
                    {
                        (context) => {
                            return <button className={styles.signOutButton} onClick={context.signOut}>
                                Выйти из аккаунта
                            </button>;
                        }
                    }
                </AuthContext.Consumer>
            </div>
            <div className={styles.modalFooter}>
                <button className={styles.backButton} onClick={() => {navigate(-1);dispatch(selectPage({page: 'userPage', state: true}));}}>
                    Назад
                </button>
            </div>
        </div>
    );
}

export default Settings;