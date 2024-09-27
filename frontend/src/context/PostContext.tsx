import { createContext } from 'react';
import ApiData from '../global_types/apiData';

type PostContextData = {
    apiData: ApiData[];
    requiredCategory: string;
    setRequiredCategory: (category: string) => void;
};

const defaultPostContextData: PostContextData = {
    apiData: [],
    requiredCategory: 'all',
    setRequiredCategory: () => { },
};

const PostContext = createContext<PostContextData>(defaultPostContextData);
export default PostContext;
