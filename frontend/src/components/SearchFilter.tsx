import { useEffect, useState } from 'react';
import ApiData from '../global_types/apiData';
import Fuse from 'fuse.js';
type SearchFilterProp = {
    unfilteredApiData:ApiData[]; 
    filteredApiData: ApiData[];
    setFilteredApiData: (data: ApiData[]) => void;
};
// this is an example of data returned by fuzzy finder
const example = [
    {
        item: {
            ID: 52,
            Title: 'Pratt Parsers: Expression Parsing Made Easy',
            Url: 'https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/',
            Body: 'A much more hands on article talking about the actual implementation of a pratt parser for a dummy language.',
            Date: '2024-09-17T14:10:43.40191Z',
            Category: 'Parsing',
        },
        refIndex: 2,
    },
    {
        item: {
            ID: 55,
            Title: 'Pratt Parsers: Expression Parsing Made Easy',
            Url: 'https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/',
            Body: 'A much more hands on article talking about the actual implementation of a pratt parser for a dummy language.',
            Date: '2024-09-17T14:10:43.40191Z',
            Category: 'Parsing',
        },
        refIndex: 4,
    },
    {
        item: {
            ID: 51,
            Title: 'Simple but Powerful Pratt Parsing',
            Url: 'https://matklad.github.io/2020/04/13/simple-but-powerful-pratt-parsing.html',
            Body: 'Helped me a bit in understanding more about pratt parsers and recursive descent parsers for my go interpreter.',
            Date: '2024-09-17T07:38:55.041995Z',
            Category: 'Parsing',
        },
        refIndex: 3,
    },
    {
        item: {
            ID: 56,
            Title: 'Simple but Powerful Pratt Parsing',
            Url: 'https://matklad.github.io/2020/04/13/simple-but-powerful-pratt-parsing.html',
            Body: 'Helped me a bit in understanding more about pratt parsers and recursive descent parsers for my go interpreter.',
            Date: '2024-09-17T07:38:55.041995Z',
            Category: 'Parsing',
        },
        refIndex: 5,
    },
];
function SearchFilter({ filteredApiData, setFilteredApiData, unfilteredApiData }: SearchFilterProp) {
    const [inputText, setInputText] = useState('');

    useEffect(() => {
        if (!inputText){
            setFilteredApiData(unfilteredApiData)
            return;
        }
        const fuseOptions = {
            // isCaseSensitive: false,
            // includeScore: false,
            // shouldSort: true,
            // includeMatches: false,
            // findAllMatches: false,
            // minMatchCharLength: 1,
            // location: 0,
            // threshold: 0.6,
            // distance: 100,
            // useExtendedSearch: false,
            // ignoreLocation: false,
            // ignoreFieldNorm: false,
            // fieldNormWeight: 1,
            keys: ['Title', 'Body'],
        };
        const fuse = new Fuse(filteredApiData, fuseOptions);
        let res = fuse.search(inputText);
        console.log('filetered data' + JSON.stringify(res));

        // Need to convert that into readable format
        let final_data = res.map((obj) => {
            return obj.item;
        });
        console.log('this is final ' + JSON.stringify(final_data));
        setFilteredApiData(final_data)
    },[inputText]);

    return (
        <div>
            <input
                type="text"
                value={inputText}
                onChange={(e) => setInputText(e.target.value)}
                placeholder="Search articles..."
                style={{ padding: '8px', width: '100%' }}
            />
        </div>
    );
}
export default SearchFilter;
