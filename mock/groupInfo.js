/*
A group has:

1. basic information: 
  i. owner: a owner
  ii. group_name
  iii. groups_avatar: group photo
  iv. description
  v. createdAt: create time
  vi. group id

2. contents
  i.  group members, # of members
  ii. posts, # post number

3. management: 
  i. delete member
  ii. delete posts
  iii. delete group
*/



const groups_owner = [
  'Soraka',
  'Ashe',
  'Sion',
  'Riven',
  'Ahri',
];

const groups_name = ['Support', 'Ad', 'Tank', 'Soldier', 'Mage'];

const descriptions = [
  'I will protect you.',
  'Kill them all!',
  'Come on hit me you little bastard.',
  'I will never fall.',
  'Let me show you some tricks.',
];

const groups_avatar = [
  
];


const groups_member =[
  [
    'Sona',
    'Soraka',
  ],
  [
    'Ashe',
    'Miss Fortune',
  ],
  [
    'Sion',
    'Sheep',
  ],
  [
    'Riven',
    'Fiona',
  ],
  [
    'Ahri',
    'Raze',
  ],
];

const titles = [
  'I hold a bow of true ice. I hold my heart.',
  'Heavy winds approaching!',
  'How noble.',
  "Don't you trust me?",
  "You'll see more with your eyes closed.",
  "Let's light it up",
  'Justice is never blind!',
  'Tricky tricky! You got the wrong Neeko!',
  'This is no ordinary instrument. More like an old friend.',
  'To heal and protect.',
];

const titles_href = [
  'https://www.leagueoflegends.com/en-us/champions/ashe/',
  'https://www.leagueoflegends.com/en-us/champions/janna/',
  'https://www.leagueoflegends.com/en-us/champions/karma/',
  'https://www.leagueoflegends.com/en-us/champions/ahri',
  'https://www.leagueoflegends.com/en-us/champions/Lulu',
  'https://www.leagueoflegends.com/en-us/champions/Lux',
  'https://www.leagueoflegends.com/en-us/champions/Morgana',
  'https://www.leagueoflegends.com/en-us/champions/Neeko',
  'https://www.leagueoflegends.com/en-us/champions/Sona',
  'https://www.leagueoflegends.com/en-us/champions/Soraka',
];

const users = [
  'Ashe',
  'Janna',
  'Karma',
  'Ahri',
  'Lulu',
  'Lux',
  'Morgana',
  'Neeko',
  'Sona',
  'Soraka',
];

const contents = [
  "Iceborn warmother of the Avarosan tribe, Ashe commands the most populous horde in the north. Stoic, intelligent, and idealistic, yet uncomfortable with her role as leader, she taps into the ancestral magics of her lineage to wield a bow of True Ice. With her people's belief that she is the mythological hero Avarosa reincarnated, Ashe hopes to unify the Freljord once more by retaking their ancient, tribal lands.",
  "Armed with the power of Runeterra's gales, Janna is a mysterious, elemental wind spirit who protects the dispossessed of Zaun. Some believe she was brought into existence by the pleas of Runeterra's sailors who prayed for fair winds as they navigated treacherous waters and braved rough tempests. Her favor and protection has since been called into the depths of Zaun, where Janna has become a beacon of hope to those in need. No one knows where or when she will appear, but more often than not, she's come to help.",
  'No mortal exemplifies the spiritual traditions of Ionia more than Karma. She is the living embodiment of an ancient soul reincarnated countless times, carrying all her accumulated memories into each new life, and blessed with power that few can comprehend. She has done her best to guide her people in recent times of crisis, though she knows that peace and harmony may come only at a considerable cost—both to her, and to the land she holds most dear.',
  'Innately connected to the latent power of Runeterra, Ahri is a vastaya who can reshape magic into orbs of raw energy. She revels in toying with her prey by manipulating their emotions before devouring their life essence. Despite her predatory nature, Ahri retains a sense of empathy as she receives flashes of memory from each soul she consumes.',
  'The yordle mage Lulu is known for conjuring dreamlike illusions and fanciful creatures as she roams Runeterra with her fairy companion Pix. Lulu shapes reality on a whim, warping the fabric of the world, and what she views as the constraints of this mundane, physical realm. While others might consider her magic at best unnatural, and at worst dangerous, she believes everyone could use a touch of enchantment.',
  "Luxanna Crownguard hails from Demacia, an insular realm where magical abilities are viewed with fear and suspicion. Able to bend light to her will, she grew up dreading discovery and exile, and was forced to keep her power secret, in order to preserve her family's noble status. Nonetheless, Lux's optimism and resilience have led her to embrace her unique talents, and she now covertly wields them in service of her homeland.",
  'Conflicted between her celestial and mortal natures, Morgana bound her wings to embrace humanity, and inflicts her pain and bitterness upon the dishonest and the corrupt. She rejects laws and traditions she believes are unjust, and fights for truth from the shadows of Demacia—even as others seek to repress it—by casting shields and chains of dark fire. More than anything else, Morgana truly believes that even the banished and outcast may one day rise again.',
  'Hailing from a long lost tribe of vastaya, Neeko can blend into any crowd by borrowing the appearances of others, even absorbing something of their emotional state to tell friend from foe in an instant. No one is ever sure where—or who—Neeko might be, but those who intend to do her harm will soon witness her true colors revealed, and feel the full power of her primordial spirit magic unleashed upon them.',
  "Sona is Demacia's foremost virtuoso of the stringed etwahl, speaking only through her graceful chords and vibrant arias. This genteel manner has endeared her to the highborn, though others suspect her spellbinding melodies to actually emanate magic—a Demacian taboo. Silent to outsiders but somehow understood by close companions, Sona plucks her harmonies not only to soothe injured allies, but also to strike down unsuspecting enemies.",
  "A wanderer from the celestial dimensions beyond Mount Targon, Soraka gave up her immortality to protect the mortal races from their own more violent instincts. She endeavors to spread the virtues of compassion and mercy to everyone she meets—even healing those who would wish harm upon her. And, for all Soraka has seen of this world's struggles, she still believes the people of Runeterra have yet to reach their full potential.",
];

const users_href = [
  'https://www.leagueoflegends.com/en-us/champions/ashe/',
  'https://www.leagueoflegends.com/en-us/champions/janna/',
  'https://www.leagueoflegends.com/en-us/champions/karma/',
  'https://www.leagueoflegends.com/en-us/champions/ahri',
  'https://www.leagueoflegends.com/en-us/champions/Lulu',
  'https://www.leagueoflegends.com/en-us/champions/Lux',
  'https://www.leagueoflegends.com/en-us/champions/Morgana',
  'https://www.leagueoflegends.com/en-us/champions/Neeko',
  'https://www.leagueoflegends.com/en-us/champions/Sona',
  'https://www.leagueoflegends.com/en-us/champions/Soraka',
];

const groups = 'support';

const groups_href = 'https://www.leagueoflegends.com/en-us/champions/';

const avatars = [
  '/heroes/Ashe_0.jpeg',
  '/heroes/Janna_0.jpeg',
  '/heroes/Karma_0.jpeg',
  '/heroes/Ahri_0.jpeg',
  '/heroes/Lulu_0.jpeg',
  '/heroes/Lux_0.jpeg',
  '/heroes/Morgana_0.jpeg',
  '/heroes/Neeko_0.jpeg',
  '/heroes/Sona_0.jpeg',
  '/heroes/Soraka_0.jpeg',
];

const postLists = [];
for(let i=0; i<5; i++) {
  const post = [];
  for(let k=0; k<2; k++) {
    post.push({
      id: i,
      owner: users[i % 10],
      owner_href: users_href[i % 10],
      title: titles[i % 10],
      title_href: titles_href[i % 10],
      logo: avatars[i % 10],
      group: groups,
      group_href: groups_href,
      updatedAt: new Date(new Date().getTime() - 1000 * 60 * 60 * 2 * i).getTime(),
      //createdAt: new Date(new Date().getTime() - 1000 * 60 * 60 * 2 * i).getTime(),
      collection: Math.ceil(Math.random() * 100) + 100,
      like: Math.ceil(Math.random() * 100) + 100,
      reply: Math.ceil(Math.random() * 10) + 10,
      content: contents[i % 10],
    });
  }
  postLists.push(post);
}


function groupInfo(userName) {
  const lists = [];
  const count = 5; //suppose has 5 groups
  for(let i = 0; i< count; i += 1) {
    lists.push({
      id: i,
      groupName: groups_name[i],
      groupDescription: descriptions[i],
      group_href: 'https://www.leagueoflegends.com/en-us/',
      //logo
      groupAvatar: avatars[i],
      createdAt: new Date(new Date().getTime() - 1000 * 60 * 60 * 2 * i).getTime(),
      groupMember: groups_member[i],
      postList: postLists[i],
    });
  }
  return lists;
}

function getGroup(req, res) {
  const params = req.query;
  const userName = params.userName;
  const result = groupInfo(userName);
  return res.json({
    data: {
      list: result,
    },
  });
}

function createdGroup(userName) {
  const count = 5;
  const lists = [];
  for(let i=0; i<count; i++) {
    lists.push({
      id: i+1,
      groupName: groups_name[i],
      groupDescription: descriptions[i],
      group_href: groups_href,
      groupAvatar: avatars[i],
      createdAt: new Date(new Date().getTime() - 1000 * 60 * 60 * 2 * i).getTime(),
      numberOfMember:Math.ceil(Math.random() * 100) + 100,
      numberOfPost:Math.ceil(Math.random() * 100) + 100,
    });
  }
  return lists;
}

function getCreated(req, res) {
  const params = req.query;
  const userName = params.userName;
  console.log(userName);
  const result = createdGroup(userName);
  return res.json({
    data: {
      list: result,
    },
  });
}

export default {
  'GET /api/getGroupInfo': getGroup,
  'GET /api/getCreatedGroup': getCreated,
  'GET /api/getJoinedGroup': getCreated,
};
