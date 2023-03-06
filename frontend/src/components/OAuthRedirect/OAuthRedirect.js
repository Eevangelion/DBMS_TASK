import { useNavigate } from "react-router-dom";
import { useGetGitQuery } from "../../services/auth"
import { getCode } from "../../store/actions/auth";
import LoadingModal from "../LoadingModal/LoadingModal";

const OAuthRedirect = () => {
    const navigate = useNavigate();
    const code = getCode();
    const {
        data: tokens,
        isLoading: loadingCode
    } = useGetGitQuery(code);
    if (loadingCode) {
        return <LoadingModal />;
    }
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
    navigate(`/feed/`);
    return;
}

export default OAuthRedirect;