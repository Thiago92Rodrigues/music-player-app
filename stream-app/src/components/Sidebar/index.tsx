import React from 'react';

// styles
import './styles.scss';

// icons
import { ReactComponent as IconHome } from '../../assets/icons/icon-home.svg';
import { ReactComponent as IconGuitar } from '../../assets/icons/icon-guitar.svg';
import { ReactComponent as IconLibrary } from '../../assets/icons/icon-library.svg';
import { ReactComponent as IconRecord } from '../../assets/icons/icon-record.svg';
import { ReactComponent as IconSearch } from '../../assets/icons/icon-search.svg';
import { ReactComponent as IconSong } from '../../assets/icons/icon-song.svg';
import { ReactComponent as IconUser } from '../../assets/icons/icon-user.svg';

export const Sidebar: React.FC = () => {
  return (
    <aside className='sidebar'>
      <ul className='sidebar__main'>
        <li>
          <a href='#'>
            <IconHome />
            <span>Home</span>
          </a>
        </li>
        <li>
          <a href='#'>
            <IconSearch />
            <span>Search</span>
          </a>
        </li>
        <li>
          <a href='#'>
            <IconUser />
            <span>Profile</span>
          </a>
        </li>
        <li>
          <a href='#'>
            <IconLibrary />
            <span>Library</span>
          </a>
        </li>
      </ul>

      <ul className='sidebar__music'>
        <h2 className='sidebar__header'>Your Music</h2>
        <li>
          <a href='#'>
            <IconSong />
            <span>Songs</span>
          </a>
        </li>
        <li>
          <a href='#'>
            <IconRecord />
            <span>Albums</span>
          </a>
        </li>
        <li>
          <a href='#'>
            <IconGuitar />
            <span>Artists</span>
          </a>
        </li>
      </ul>

      <ul className='sidebar__playlists'>
        <h2 className='sidebar__header'>Playlists</h2>
      </ul>
    </aside>
  );
};
