import { useNavigate } from "react-router-dom";
import { useGetGitQuery } from "../../services/auth"
import { getCode } from "../../store/actions/auth";

const OAuthRedirect = () => {
    const code = getCode();
    const navigate = useNavigate();
    const {
        data: user,
        isLoading: loadingCode
    } = useGetGitQuery(code);
    if (loadingCode) {
        return <></>;
    }
    localStorage.setItem("userID", user.id);
    localStorage.setItem("userName", user.name);
    localStorage.setItem("userRole", user.role);
    localStorage.setItem("access_token", user.token);
    navigate('/feed/');
}

export default OAuthRedirect;