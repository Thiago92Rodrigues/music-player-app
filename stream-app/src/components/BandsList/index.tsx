import React, { useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';

import { api, staticFilesAddress } from '../../services/api';
import { Artist as Band } from '../../types';

// styles
import './styles.scss';

export const BandsList: React.FC = () => {
  const history = useHistory();

  const [bands, setBands] = useState<Band[]>([]);

  useEffect(() => {
    api.get('/artists').then(response => setBands(response.data));
  }, []);

  function handleBandClick(id: string) {
    console.log('Band click ', id);
    history.push(`/artists/${id}`);
  }

  return (
    <div className='bands'>
      <div className='bands__header'>
        <h2>Bands</h2>
      </div>

      <ul className='bands__list'>
        {bands.map(band => (
          <li
            className='bands__list__item'
            key={band.id}
            onClick={() => handleBandClick(band.id)}
            style={{
              backgroundImage: `url(${staticFilesAddress}/files/?file=${band.profile_img}`,
            }}
          >
            <div className='bands__list__item__info'>
              <span className='band__name'>{band.name}</span>
              <span className='band__genre'>{band.genre}</span>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};
