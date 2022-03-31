import { LikeOutlined, LoadingOutlined, MessageOutlined, StarOutlined } from '@ant-design/icons';
import { Button, Card, Col, Form, List, Row, Select, Tag } from 'antd';
import React, { useState } from 'react';
import { useRequest } from 'umi';
import ArticleListContent from './components/ArticleListContent';
import StandardFormRow from './components/StandardFormRow';
import TagSelect from './components/TagSelect';
import { queryList } from '@/services/getList';
import styles from './style.less';

const { Option } = Select;
const FormItem = Form.Item;
const pageSize = 20;
const pageNumber = 1;

const waitTime = (time = 100) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(true);
    }, time);
  });
};

const Articles = () => {
  const [form] = Form.useForm();
  const { data, reload, loading, loadMore, loadingMore } = useRequest(
    async() => {
      const result = await queryList({
        PageNO: pageNumber,
        PageSize: pageSize,
      });
      console.log(result);
      return result;
    },
    {
      formatResult: result => result,
      loadMore: true,
    }    
  );

  // async function getArticleList() {
  //   console.log("getting data about articles")
  //   let result = []
  //   const {data} =  await fetch(
  //     '/api/article/getarticlelist?PageNO=1&PageSize=20',
  //     {
  //       method: 'GET',
  //       credentials: 'include'
  //     }
  //   ).then(response => response.json())
  //   console.log(data)
  // }
  // console.log("====== getting async")
  // getArticleList()
  // console.log("====== data")
  
//  response.then((data) => {
//    console.log(data)
//  })

  console.log(data);
  const list = [];
  

  if(typeof(data[0])!='undefined') {
    var size = Object.keys(data).length;

    for(let i=0; i<size-1; i++) {
      console.log(data[i]);
      list.push(data[i]);
    }
  }
  
  // for(let i=0; i<10; i++) {
  //   console.log(data[i]);
  //   list.push(data[i]);
  // }
 console.log(list);


  const onCCollection = async(values) => {
    console.log(values);
    if(values.type === 'star-o') {
      if(values.value === '1') {
        return (
          <IconText key="collection" type="star-o" value="0" text={values.text--} />
        );
      }
      else {
        return (
          <IconText key="collection" type="star-o" value="1" text={item.collection} />
        )
      }
    }
  }

  const IconText = ({ type, text }) => {
    switch (type) {
      case 'star-o':
        return (
          <span>
            <StarOutlined
              style={{
                marginRight: 8,
              }}
            />
            {text}
          </span>
        );

      case 'like-o':
        return (
          <span>
            <LikeOutlined
              style={{
                marginRight: 8,
              }}
            />
            {text}
          </span>
        );

      case 'message':
        return (
          <span>
            <MessageOutlined
              style={{
                marginRight: 8,
              }}
            />
            {text}
          </span>
        );

      default:
        return null;
    }
  };

  const formItemLayout = {
    wrapperCol: {
      xs: {
        span: 24,
      },
      sm: {
        span: 24,
      },
      md: {
        span: 12,
      },
    },
  };

  const onCollection = async(values) => {
    console.log(values);
    let count = values;
    count++;
    return count;
  }

  const loadMoreDom = list.length > 0 && (
    <div
      style={{
        textAlign: 'center',
        marginTop: 16,
      }}
    >
      <Button
        onClick={loadMore}
        style={{
          paddingLeft: 48,
          paddingRight: 48,
        }}
      >
        {loadingMore ? (
          <span>
            <LoadingOutlined /> Loading...
          </span>
        ) : (
          'Load More'
        )}
      </Button>
    </div>
  );

  return (
    <>
      <Card bordered={false}>
        <Form
          layout="inline"
          form={form}
          initialValues={
            {
              //owner: ['wjh', 'zxx'],
            }
          }
          onValuesChange={reload}
        >
        <TagSelect expandable>
          <TagSelect.Option value="cat1">Sports</TagSelect.Option>
          <TagSelect.Option value="cat2">Professors</TagSelect.Option>
          <TagSelect.Option value="cat3">Courses</TagSelect.Option>
          <TagSelect.Option value="cat4">Daily Life</TagSelect.Option>
          <TagSelect.Option value="cat5">Movies</TagSelect.Option>
        </TagSelect>
        </Form>
      </Card>
      <Card
        style={{
          marginTop: 10,
        }}
        bordered={false}
        bodyStyle={{
          padding: '8px 32px 32px 32px',
        }}
      >
        <List
          size="large"
          loading={loading}
          rowKey="id"
          itemLayout="vertical"
          loadMore={loadMoreDom}
          dataSource={list}
          renderItem={(item) => (
            <List.Item
              key={item.ID}
              actions={[
                <IconText key="collection" type="star-o" text={item.collection}  />,
                <IconText key="like" type="like-o" text={item.like} />,
                <IconText key="reply" type="message" text={item.reply} />,
              ]}
            >
              <List.Item.Meta
                title={
                  <a className={styles.listItemMetaTitle} href={"/group/post?"+item.ID}>
                    {item.Title}
                  </a>
                }
              />
              <ArticleListContent data={item} />
            </List.Item>
          )}
        />
      </Card>
    </>
  );
};

export default Articles;