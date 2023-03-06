import LoadingModal from '../components/LoadingModal/LoadingModal';
import { useGetUnbanDateQuery } from '../services/service';
import '../styles/BannedPage.css';

const BannedPage = (props) => {
    const {
        data: unparsedUnbanDate,
        isLoading
    } = useGetUnbanDateQuery();

    if (isLoading) {
        return <LoadingModal />;
    }
    const unbanDate = unparsedUnbanDate.split('T')[0];
    return (
        <div className="banned-page">
            <div className="banned-text">
                <h1>На вас поступила жалоба и администрация приняла решение заблокировать ваш аккаунт</h1>
                <br/>
                <h2>Блокировка продлится до {unbanDate}</h2>
            </div>
        </div>
    );
};

export default BannedPage;