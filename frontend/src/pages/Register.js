import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useRegisterUserMutation } from "../services/auth";
import styles from "../styles/Register.module.css";

const RegisterPage = () => {
    const navigate = useNavigate();
    const [emailText, setEmailText] = useState('');
    const [usernameText, setUsernameText] = useState('');
    const [passwordText, setPasswordText] = useState('');
    const [registerUser] = useRegisterUserMutation();

    const handleRegister = (email, name, password) => {
        registerUser({email: email, username: name, password: password})
        navigate("/login");
    };
    return (
        <div className={styles.modalWindow}>
            <div className={styles.modalHeader}>
                Регистрация
            </div>
            <div className={styles.modalBody}>
                <div className={styles.emailForm}>
                    <p>Почта</p>
                    <div className={styles.emailField}>
                        <input   className={styles.signinEmail} 
                                    placeholder="Введите почту" 
                                    onChange={e=>setEmailText(e.target.value)} 
                                    value={emailText} 
                                    required={true}
                        >            
                        </input>
                    </div>
                </div>
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
            </div>
            <div className={styles.modalFooter}>
                <button 
                    className={styles.registerButton}
                    onClick={(event) => {
                        (emailText && usernameText && passwordText) ? 
                        handleRegister(emailText, usernameText, passwordText) : 
                        event.preventDefault()
                    }}
                >
                    Зарегистрироваться
                </button>
            </div>
        </div>
    );
}

export default RegisterPage;