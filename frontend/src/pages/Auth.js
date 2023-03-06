import { useState } from "react";
import { useDispatch } from "react-redux";
import { useNavigate, Link } from "react-router-dom";
import { useLoginUserMutation } from "../services/auth";
import { selectPage } from "../store/reducers/page";
import styles from "../styles/Auth.module.css";

const clientID = process.env.REACT_APP_CLIENT_ID;

const AuthPage = () => {
    const handleClick = () => {
        window.location.href = `https://github.com/login/oauth/authorize?client_id=${clientID}`;
    }
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [usernameText, setUsernameText] = useState('');
    const [passwordText, setPasswordText] = useState('');
    const [loginUser] = useLoginUserMutation();
    localStorage.clear();
    const handleLogin = async (name, password) => {
        await loginUser({username: name, password: password}).then((response) => {
            const tokens = response.data;
            const accessToken = tokens.jwt_token;
            const refreshToken = tokens.refresh_token;
            const base64Url = accessToken.split('.')[1];
            const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
            const jsonPayload = decodeURIComponent(window.atob(base64).split('').map((c) => {
                return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
            }).join(''));
            const data = JSON.parse(jsonPayload);
            localStorage.setItem("userID", data.user_id);
            localStorage.setItem("userName", data.username);
            localStorage.setItem("userRole", data.role);
            localStorage.setItem("access_token", accessToken);
            localStorage.setItem("token_exp_time", data.exp);
            localStorage.setItem("refresh_token", refreshToken);
        })
        dispatch(selectPage({page: 'userPage', state: true}));
        dispatch(selectPage({page: 'feed', state: true}));
        dispatch(selectPage({page: 'searchPage', state: true}));
        dispatch(selectPage({page: 'subscribes', state: true}));
        navigate("/feed/");
    };
    return (
        <div className={styles.modalWindow}>
            <div className={styles.modalHeader}>
                Авторизация
            </div>
            <div className={styles.modalBody}>
                <div style={{paddingLeft: "4vw"}}>Авторизируйтесь, чтобы пользоваться сайтом</div>
                <div className={styles.usernameForm}>
                    <p>Имя</p>
                    <div className={styles.usernameField}>
                        <input   className={styles.signinUsername} 
                                    placeholder="Введите имя" 
                                    onChange={e=>setUsernameText(e.target.value)} 
                                    value={usernameText} 
                                    required={true}
                        >            
                        </input>
                    </div>
                </div>
                <div className={styles.passwordForm}>
                    <p>Пароль</p> 
                    <div className={styles.passwordField}>
                        <input   className={styles.signinPassword} 
                                    placeholder="Введите пароль" 
                                    onChange={e=>setPasswordText(e.target.value)} 
                                    value={passwordText} 
                                    required={true}
                                    type="password"
                        >
                        </input>
                    </div>
                </div>
                <button 
                    className={styles.loginButton}
                    onClick={(event) => {
                        (usernameText && passwordText) ? 
                        handleLogin(usernameText, passwordText) : 
                        event.preventDefault()
                    }}
                >
                    Авторизироваться
                </button>
                <Link   className={styles.registerLink}
                        to={`/register/`}
                >
                    Нет учётной записи? Зарегистрируйтесь!
                </Link>
            </div>
            <div className={styles.modalFooter}>
                <button className={styles.loginGitButton}
                        onClick={handleClick}>
                    Авторизироваться с помощью Github
                </button>
            </div>
        </div>
    );
}

export default AuthPage;