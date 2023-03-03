import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { useLoginUserMutation } from "../services/auth";
import styles from "../styles/Auth.module.css";

const clientID = process.env.REACT_APP_CLIENT_ID;

const AuthPage = () => {
    const handleClick = () => {
        window.location.href = `https://github.com/login/oauth/authorize?client_id=${clientID}`;
    }
    const navigate = useNavigate();
    const [usernameText, setUsernameText] = useState('');
    const [passwordText, setPasswordText] = useState('');
    const [loginUser] = useLoginUserMutation();

    const handleLogin = (name, password) => {
        loginUser({username: name, password: password}).then((response) => {
            const tokens = response.data;
            const accessToken = tokens.jwt_token;
            const refreshToken = tokens.refresh_token;
            const base64Url = accessToken.split('.')[1];
            const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
            const jsonPayload = decodeURIComponent(window.atob(base64).split('').map((c) => {
                return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
            }).join(''));
            const user = JSON.parse(jsonPayload);
            localStorage.setItem("userID", user.user_id);
            localStorage.setItem("userName", user.username);
            localStorage.setItem("userRole", user.role);
            localStorage.setItem("access_token", accessToken);
            localStorage.setItem("token_exp_time", user.exp);
            localStorage.setItem("refresh_token", refreshToken);
        })
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
                    <text>Имя</text>
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
                    <text>Пароль</text> 
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
                    onClick={() => handleLogin(usernameText, passwordText)}
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