# API分析
TIP. 从2017.12.15开始，部分接口需要设置Header(Welove-UA: [Device:PHONE_NO][OSV:7.1.1][CV:Android4.0.3][WWAN:0][zh_CN][platform:tencent][WSP:2])

### 1. URL
`access_token`: `56294XXXX343086-27152XXXXXa5746dd`(抓包获取，登录之后保持不变)

`sig`: `WXjA+ujXTVKUfv9lfVMGo6pxbis=` (**加密字段**)

**加密方式**:`{GET|POST}&{url}&{content}`并且`url`和`content`都是经过`Base64`编码，
`POST`、`url` 都是固定的，`content`则是请求的data信息出除去`sig`字段的所有字段按照首字母ascii升序排序，其间用&分隔，对这一构造得到的字符串进行`HmacSHA1`加密之后再进行`Base64`编码就是请求中需要的`sig`了[查看详细](http://blog.csdn.net/hanziyuan08/article/details/52606908)

### 2. 我们的家
```
task_type: 
{
    休息: 7,
    洗澡: 6,
    吃饭: 4,
    睡觉: 5,
    互动: 11
}

获取love_space_id: 
    http://api.welove520.com/v5/useremotion/getone
参数:
    access_token=562949961343086-2ca7e299a09974dd0&app_key=ac5f34563a4344c4&user_id=0&sig=n7DSsDmFBjU8tWnJXINN8iN1GNo=
返回值:
    {
        "result": 1,
        "love_space_id": 844424932415867,
        "emotion_last": 8,
        "emotion_cur": 0,
        "user_id": 562949961343055,
        "set_time": 1474348805820
    }

随机拜访：
    http://api.welove520.com/v1/game/house/info
参数：
    access_token=56294996134XXXXXXXX99a09974dd0&love_space_id=random&sig=EqX50TBS1oaCkDKmmA7Y2Jp9SKI=
返回值：
    {
        "result": 1,
        "messages": [
            {
                "view_ad_count": 0,
                "users": [
                    {
                        "gender": 1,
                        "user_id": "562949973509257",
                        "gender_update": 1
                    },
                    {
                        "gender": 0,
                        "user_id": "562949973509232",
                        "gender_update": 1
                    }
                ],
                "faces": [
                    {
                        "face": [
                            {
                                "color": 0,
                                "status": 1,
                                "goods_id": 50600001
                            },
                            ...
                        ],
                        "user_id": "562949973509257"
                    },
                    ...
                ],
                "pay_old_goods": 0,
                "cur_house_num": 0,
                "medals": [
                    {
                        "level": 1,
                        "owner_id": "562949973509257",
                        "medal_id": 1
                    },
                    ...
                ],
                "msg_type": 30,
                "fashions": [
                    {
                        "fashion": [
                            {
                                "id": 459319,
                                "count": 1,
                                "goods_id": 20300031,
                                "wear": 1
                            },
                            ...
                        ],
                        "user_id": "562949973509257"
                    },
                    ...
                ],
                "house": {
                    "estate_count": 1,
                    "is_stroke_pet": 0,
                    "headurl": "",
                    "love_space_id": "8444xxxxxxxxx2707",
                    "name": "",
                    "like_today": 1,
                    "house_id": 6690739,
                    "photo_id": "0",
                    "cover_id": 0
                },
                "decorations": [
                    {
                        "id": 2285258,
                        "decorate": 1,
                        "status": 1,
                        "rotation": 90,
                        "goods_id": 11900107,
                        "z": 0,
                        "house_num": 0,
                        "y": 6,
                        "parent_id": 0,
                        "x": 15
                    },
                    ...
                ],
                "game_version": 0
            }
        ]
    }
    
称赞:
    http://api.welove520.com/v1/game/house/task
参数:
    task_type=8&house_num=0&access_token=56294996134XXXXXXXX99a09974dd0&love_space_id=8444xxxxxxxxx2707&sig=Ca8fn0iQ7attLlZjZBYUdY9IXPQ=
返回值:
    {
        "result": 1,
        "messages": [
            {
                "lasting_days": 0,
                "is_home_done": 1,
                "count": 1,
                "msg_type": 37,
                "task_type": 8,
                "remain_time": 0
            },
            {
                "msg_type": 31,
                "house": {
                    "love_space_id": "844424932415867",
                    "recycle_card": 0,
                    "diamond": 13371,
                    "house_id": 2283899
                }
            }
        ]
    }
```

### 3. 爱情树
```
获取爱情树信息：
    http://api.welove520.com/v1/game/tree/getInfo?access_token=562949961313211-2cxxxxx7e299a0997xxx0&app_key=ac5XXXXXa4344c4&screen_type=102&tree_version=30&sig=XLqfKUjNrU0PlpKeFOA4SwroVRs=
返回值：
{
    "result": 1,
    "next_level_growth": 1100,
    "lover_lack_sunlight": 1,
    "new_op_record": 1,
    "stage": 2,
    "next_stage_growth": 40,
    "level": 31,
    "name": "美美的割割",
    "level_growth": 510,
    "card": 0,
    "gold": 160,
    "lack_water": 1,
    "lover_lack_water": 1,
    "lack_sunlight": 1,
    "decoration": [
        {
            "position": 4,
            "goods_id": 101
        },
        ...
    ],
    "tree_id": 844424932415867,
    "growth": 24635
}

浇水：
http://api.welove520.com/v1/game/tree/op
请求值：
access_token=562949961313211-2cxxxxx7e299a0997xxx0&app_key=ac5XXXXXa4344c4&op=1&sig=2MH+eM+A4deR58Zimzk5Tk1Zbkk=
返回值：
{
    "extra_growth": 0,
    "result": 1,
    "next_level_growth": 1100,
    "stage": 2,
    "next_stage_growth": 35,
    "level": 31,
    "name": "美美的割割",
    "level_growth": 515,
    "card": 0,
    "lover_lasting_days": 915,
    "tree_id": 844424932415867,
    "lasting_days": 915,
    "level_up": 0,
    "complete_today": 0,
    "growth_increase": 5,
    "growth": 24640
}

晒太阳：
http://api.welove520.com/v1/game/tree/op
请求值：
access_token=562949961313211-2cxxxxx7e299a0997xxx0&app_key=ac5XXXXXa4344c4&op=2&sig=MpxK/hUWRSnyDFkpSnOl06d3UiQ=
返回值：
{
    "extra_growth": 0,
    "result": 1,
    "next_level_growth": 1100,
    "stage": 2,
    "next_stage_growth": 30,
    "level": 31,
    "name": "美美的割割",
    "level_growth": 520,
    "card": 0,
    "lover_lasting_days": 915,
    "tree_id": 844424932415867,
    "lasting_days": 915,
    "level_up": 0,
    "complete_today": 1,
    "growth_increase": 5,
    "growth": 24645
}

爱情树记录查询：
http://api.welove520.com/v1/game/tree/records?access_token=562949961313211-2cxxxxx7e299a0997xxx0&app_key=ac5XXXXXa4344c4&sig=1nuLuqYjMhOcU2xlo7pG0Mz1n00=
返回值：
{
    "result": 1,
    "lasting_days": 915,
    "lover_lack_water": 1,
    "cur_month": 9,
    "lover_lack_sunlight": 1,
    "card_balance": 11,
    "low_price_left": 1,
    "lover_records": [
        {
            "lasting_days": 0,
            "invalid": 0,
            "complete": 0,
            "date": "2016-09-15"
        },
        ...
    ],
    "low_price_total": 1,
    "records": [
        {
            "lasting_days": 915,
            "invalid": 0,
            "complete": 1,
            "date": "2016-09-15"
        },
        ...
    ],
    "lover_lasting_days": 274,
    "tree_id": 844424932415867
}
```

### 4.宠物
```
宠物状态:
    http://api.welove520.com/v1/game/house/pet/task/list
参数:
    access_token=562949961343086-210646ecc6c6d9483e&sig=RTIJqvCeipVTZOsfhN9GWYN7%2BjA=
返回值:
    {
        "result": 1,
        "messages": [
            {
                "msg_type": 110,
                "pets": [
                    {
                        "pet_id": 20203,
                        "pet_tasks": [
                            {
                                "count": 1,
                                "task_type": 1,
                                "remain_time": 0 //剩余时间
                            },
                            {
                                "count": 1,
                                "task_type": 2,
                                "remain_time": 0
                            },
                            {
                                "count": 1,
                                "task_type": 3,
                                "remain_time": 0
                            },
                            {
                                "count": 1,
                                "task_type": 4,
                                "remain_time": 0
                            }
                        ]
                    }
                ]
            },
            {
                "count": 0,
                "msg_type": 136
            }
        ]
    }
    
抚摸:
    http://api.welove520.com/v1/game/house/pet/task/do
参数:
    access_token=562949961343086-210646ecc6c6d9483e&pet_id=20203&task_type=4&sig=G2xPwEuSqJwzx7Jn83zvQMlwlh0=
    task_type(4抚摸 1吃饭 2喝水 3洗澡)
返回值:
    {
        "result": 1,
        "messages": [
            {
                "pet_id": 20203,
                "count": 1,
                "msg_type": 109,
                "task_type": 4,
                "remain_time": 304
            },
            {
                "pet_id": 20203,
                "inc_exp": 0,
                "msg_type": 114
            },
            {
                "pet_id": 20203,
                "time": "2016-05-14 11:56:17",
                "level": 25,
                "msg_type": 106,
                "name": "泡芙",
                "thirsty_time_left": 11311,
                "dirty_time_left": 32714,
                "hungry_time_left": 11303,
                "growth": 6590
            },
            {
                "pet_id": 20203,
                "msg_type": 142,
                "tasks": [
                    {
                        "count": 3,
                        "status": 0,
                        "task_id": 1
                    },
                    {
                        "count": 1,
                        "status": 0,
                        "task_id": 2
                    },
                    {
                        "count": 2,
                        "status": 2,
                        "task_id": 3
                    }
                ]
            }
        ]
    }
```

### 5. 农场 [物品ID](https://github.com/Bpazy/wl520/blob/master/document/ItemID.md)
```
概况:
    http://api.welove520.com/v1/game/farm/panorama
参数:
    access_token=562949961343086-xxxxxxb7bf7c154d0d&obstacle_num=1&sig=mSJtnlo4Nw5r1wDooVs%2FLXntOjo%3D
返回值:
    {
        "result": 1,
        "messages": [ //数组内每个元素代表的意义不一样
            { //农场概况，比如name一类
                "roulette_coin": 10,
                "goods_capacity": 175,
                "female_head_url": "http://welove.b0.upaiyun.com/img/20161222/2216/28514b3e.jpg_large",
                "obstacles": [
                    {
                        "rotate": 0,
                        "item_id": 110011,
                        "x": 37,
                        "y": 27,
                        "build_status": 0,
                        "time": 1484237379393,
                        "id": 42308250
                    }
                ],
                "op_time": 1485836167650,
                "factories": [
                    {
                        "rotate": 0,
                        "level": 0,
                        "item_id": 105001,
                        "x": 16,
                        "y": 9,
                        "build_status": 0,
                        "time": 1484237379393,
                        "id": 42308248,
                        "capacity": 3,
                        "products": [
                            {
                                "left_time": -1,
                                "item_id": 203003,
                                "product_id": 208782955,
                                "time": 1485833908764,
                                "produce_start_time": 1485834506309
                            }
                        ]
                    }
                ],
                "decorations": [
                    {
                        "rotate": 0,
                        "item_id": 111015,
                        "x": 20,
                        "y": 17,
                        "id": 52763009,
                        "status": 0
                    }
                ],
                "gold_coin": 8186,
                "beehives": [],
                "base_buildings": [
                    {
                        "rotate": 0,
                        "item_id": 101001,
                        "x": 16,
                        "y": 14,
                        "build_status": 0,
                        "time": 1484237379393,
                        "id": 42308229
                    }
                ],
                "unlock_blocks": [],
                "fruit_trees": [
                    {
                        "revive_count": 1,
                        "rotate": 0,
                        "revive_status": 0,
                        "item_id": 106001,
                        "product_status": 0,
                        "revive_user_id": "562949973408733",
                        "produce_start_time": 1485060688203,
                        "collect_count": 4,
                        "left_time": -717879447,
                        "revive_farm_id": "844424937994272",
                        "x": 5,
                        "y": 11,
                        "time": 1484574028755,
                        "id": 73144925
                    }
                ],
                "gift_box": {
                    "box_id": 2,
                    "cost_rainbow_coin": 0
                },
                "crops_capacity": 175,
                "male_id": "562949961343086",
                "male_guided": 1,
                "level": 23,
                "is_alpha": 0,
                "female_id": "562949961343055",
                "level_exp": 717,
                "yards": [
                    {
                        "rotate": 0,
                        "item_id": 103001,
                        "x": 20,
                        "y": 28,
                        "animals": [
                            {
                                "left_time": -1485834967651,
                                "last_collect_time": 1485834735876,
                                "product_status": 0,
                                "animal_id": 2018456,
                                "animal_item_id": 104001,
                                "last_feed_time": -1
                            }
                        ],
                        "time": 1484237379393,
                        "id": 42308247
                    }
                ],
                "flower_bushes": [],
                "rainbow_coin": 2,
                "female_guided": 1,
                "love_heart": 50,
                "male_head_url": "http://welove.b0.upaiyun.com/img/20161222/2232/12b1fcf6.jpg_large",
                "official_farm_id": "844424933116515",
                "name": "元露酱",
                "msg_type": 2,
                "fields": [
                    {
                        "left_time": 5424094,
                        "rotate": 0,
                        "plant_time": 1485832591744,
                        "item_id": 102001,
                        "x": 15,
                        "product_status": 0,
                        "y": 24,
                        "time": 1484237379393,
                        "id": 42308238,
                        "plant_item_id": 201008
                    }
                ],
                "total_exp": 54279
            },
            { //未知
                "op_time": 1485836167651,
                "msg_type": 113,
                "place": 1,
                "id": 7598425,
                "status": 0
            },
            { //未知
                "op_time": 1485836167651,
                "msg_type": 90
            },
            { //未知
                "op_time": 1485836167651,
                "msg_type": 16,
                "orders": [
                    {
                        "appear_left_time": -1142356,
                        "time": 1485835025295,
                        "slot": 1,
                        "exp": 5,
                        "order_id": 80608485,
                        "items": [
                            {
                                "item_id": 201004,
                                "count": 20
                            }
                        ],
                        "status": 1,
                        "coin": 221
                    }
                ]
            },
            { //未知
                "hire_count": 1,
                "op_time": 1485836167651,
                "msg_type": 29
            },
            { //汽车订单
                "op_time": 1485836167652,
                "msg_type": 15,
                "orders": [
                    {
                        "special": 0,
                        "time_left": -1,
                        "voucher_item_id": 0,
                        "op_time": 1485836167652,
                        "slot": 6,
                        "exp": 237,
                        "order_id": 30285736,
                        "items": [
                            {
                                "item_id": 202003,
                                "count": 1
                            }
                        ],
                        "status": 0,
                        "coin": 313,
                        "buyer": 2
                    }
                ]
            },
            { //签到信息
                "lover_records": [
                    {
                        "date": "2017-01-31",
                        "lasting_days": 1,
                        "invalid": 0,
                        "complete": 1
                    }
                ],
                "lasting_days": 10,
                "record": [
                    {
                        "date": "2017-01-31",
                        "lasting_days": 10,
                        "invalid": 0,
                        "complete": 1
                    }
                ],
                "today": "2017-01-31",
                "msg_type": 118,
                "lover_lasting_days": 1
            },
            { //自己货摊信息
                "last_free_ad_time": 0,
                "op_time": 1485836167652,
                "msg_type": 20,
                "stall_items": [ //自己货摊售卖物品
                    {
                        "item_id": 204023,
                        "count": 3,
                        "last_ad_time": 1485834677817,
                        "id": 32148667,
                        "slot": 4,
                        "status": 1,
                        "coin": 604
                    }
                ],
                "ad_auth": 0,
                "capacity": 6,
                "farm_id": "844424932415867"
            },
            { //轮船信息
                "ships": [
                    {
                        "boxes": [
                            {
                                "user_id": "562949961343055",
                                "item_id": 202003,
                                "count": 5,
                                "thank": 0,
                                "id": 7598422,
                                "slot": 1,
                                "exp": 115,
                                "status": 0,
                                "coin": 140,
                                "farm_id": "844424932415867"
                            }
                        ],
                        "leave_time_left": 31872398,
                        "score": 16,
                        "enter_time_left": 0,
                        "voucher_count": 1,
                        "position": 1,
                        "enter_time": 1485810140051,
                        "exp": 174,
                        "ship_id": 1065940,
                        "status": 1
                    }
                ],
                "op_time": 1485836167653,
                "msg_type": 38,
                "farm_id": "844424932415867"
            },
            { //未知
                "warehouses": [
                    {
                        "category": 1,
                        "items": [
                            {
                                "item_id": 201002,
                                "count": 29
                            }
                        ]
                    }
                ],
                "op_time": 1485836167653,
                "msg_type": 3
            }
        ]
    }

预览:
    http://api.welove520.com/v1/game/farm/preview
参数:
    access_token=562949961343086-xxxxxxb7bf7c154d0d&farm_id=844424933116515&sig=USuL4XorvUPcUlAwwkhSBdhTeVY%3D
返回值:
    {
        "result": 1,
        "messages": [
            {
                "preview": {
                    "male_id": "562949963376375",
                    "alive": 1,
                    "level": 56,
                    "male_head_url": "http://welove.b0.upaiyun.com/img/20161208/2233/eb1ec359.jpg_large",
                    "female_head_url": "http://welove.b0.upaiyun.com/img/20161208/2234/5f79a115.jpg_large",
                    "female_id": "562949963376390",
                    "name": "小微与小爱",
                    "farm_id": "844424933116515"
                },
                "msg_type": 120
            }
        ]
    }

同步信息:
    http://api.welove520.com/v1/game/farm/operations/sync
参数:
    access_token=562949961343086-21410ab7bf7c154d0d&sig=Qa0TDzyPnP2iqRnl7MTWvtmSheU%3D
返回值:
    {
        "result": 1,
        "messages": [
            {
                "op_time": 1485840277256,
                "msg_type": 100
            },
            {
                "user_id": 562949961343055,
                "op_time": 1485840277256,
                "msg_type": 110,
                "status": 0
            },
            {
                "need_helps": 0,
                "mails": 1,
                "achievements": 0,
                "msg_type": 116
            }
        ]
    }

打开礼物箱:
    http://api.welove520.com/v1/game/farm/giftbox/open
参数:
    access_token=562949961343086-21410ab7bf7c154d0d&rainbow_coin=0&sig=2kL0PunZWn%2FxbK7VLuPkhEyF%2FeI%3D
返回值:
    {
        "result": 1,
        "messages": [
            {
                "op_time": 1485838495984,
                "exp_inc": 156,
                "msg_type": 103,
                "farm_id": "844424932415867"
            },
            {
                "op_time": 1485838495984,
                "msg_type": 101,
                "exp": 156
            }
        ]
    }

删除礼物箱:
    http://api.welove520.com/v1/game/farm/giftbox/del
参数:
    access_token=562949961343086-21410ab7bf7c154d0d&sig=rr2m4ciQq8JgDEOXpXDCzR9S97Y%3D
返回值:
    {
        "result": 1,
        "messages": [
            {
                "op_time": 1485840335325,
                "msg_type": 121
            }
        ]
    }

查询广告牌售卖物品:
    http://api.welove520.com/v1/game/farm/ad/query
参数:
    access_token=562949961343086-21410ab7bf7c154d0d&sig=OXqTXVVia4oGv%2Fby8Rc8z4hGtxk%3D
返回值:
    {
        "result": 1,
        "messages": [
            {
                "op_time": 1485838510201,
                "msg_type": 28,
                "ad_items": [
                    {
                        "item_id": 204002,
                        "count": 2,
                        "op_time": 1485838510201,
                        "need_help": 1,
                        "seller_farm_id": "844424938221375",
                        "head_url_famale": "http://welove.b0.upaiyun.com/img/20170116/2326/e5a490a7.jpg_large",
                        "head_url_male": "http://welove.b0.upaiyun.com/img/20170116/2306/b32fbb9a.jpg_large",
                        "id": 34263655,
                        "farm_name": "L.I.Z",
                        "coin": 97
                    }
                ]
            }
        ]
    }

购买广告牌售卖物品:
    http://api.welove520.com/v1/game/farm/stall/buy
参数:
    access_token=562949961343086-21410ab7bf7c154d0d&seller_farm_id=844424935769649&sig=n4wclZvyiw6yHvDVtdZCApp6s5M%3D&stall_sale_id=34361005
返回值:
    {
        "result": 1,
        "messages": [
            {
                "stall_item": {
                    "buyer_head_url": "http://welove.b0.upaiyun.com/img/20161222/2232/12b1fcf6.jpg_large",
                    "buyer_farm_name": "元露酱",
                    "id": 34361005
                },
                "op_time": 1485843672864,
                "msg_type": 23
            },
            {
                "warehouses": [
                    {
                        "category": 2,
                        "items_inc": [
                            {
                                "item_id": 204017,
                                "count": 2
                            }
                        ]
                    }
                ],
                "op_time": 1485843672864,
                "msg_type": 102,
                "farm_id": "844424932415867"
            },
            {
                "op_time": 1485843672864,
                "gold_cost": 64,
                "msg_type": 103,
                "farm_id": "844424932415867"
            }
        ]
    }

收割庄稼:
    http://api.welove520.com/v1/game/farm/factory/harvest
参数:
    access_token=562949961343086-21410ab7bf7c154d0d&factory_id=42308249&product_ids=208600635&rainbow_coin=0&sig=HpmhB8JlranXw51XYciWDHbWJbs%3D
返回值:
    {
        "result": 1,
        "messages": [
            {
                "op_time": 1485840269431,
                "exp_inc": 8,
                "msg_type": 103,
                "farm_id": "844424932415867"
            },
            {
                "warehouses": [
                    {
                        "category": 2,
                        "items_inc": [
                            {
                                "item_id": 204002,
                                "count": 1
                            }
                        ]
                    }
                ],
                "op_time": 1485840269431,
                "msg_type": 102,
                "farm_id": "844424932415867"
            },
            {
                "factory_id": 42308249,
                "product_ids": [
                    208600635
                ],
                "op_time": 1485840269431,
                "msg_type": 10
            }
        ]
    }

种植庄稼:
    http://api.welove520.com/v1/game/farm/crops/plant
参数:
    access_token=562949961343086-21410ab7bf7c154d0d&farmlands=%5B%7B%22id%22%3A118814820%2C%22x%22%3A12%2C%22last_interval%22%3A1%2C%22y%22%3A24%7D%2C%7B%22id%22%3A67191727%2C%22x%22%3A13%2C%22last_interval%22%3A1%2C%22y%22%3A24%7D%2C%7B%22id%22%3A42308241%2C%22x%22%3A14%2C%22last_interval%22%3A1%2C%22y%22%3A25%7D%2C%7B%22id%22%3A83803299%2C%22x%22%3A14%2C%22last_interval%22%3A1%2C%22y%22%3A26%7D%2C%7B%22id%22%3A42578886%2C%22x%22%3A14%2C%22last_interval%22%3A1%2C%22y%22%3A29%7D%2C%7B%22id%22%3A73106450%2C%22x%22%3A15%2C%22last_interval%22%3A1%2C%22y%22%3A30%7D%2C%7B%22id%22%3A42578866%2C%22x%22%3A14%2C%22last_interval%22%3A1%2C%22y%22%3A24%7D%5D&item_id=201001&sig=IvxXLfaNfOK9qYvTkqO3qFXHpSI%3D
返回值:
    {
        "result": 1,
        "messages": [
            {
                "item_id": 201001,
                "op_time": 1485840313517,
                "msg_type": 5,
                "farmlands": [
                    {
                        "left_time": 119999,
                        "x": 12,
                        "y": 24,
                        "id": 118814820
                    }
                ]
            },
            {
                "warehouses": [
                    {
                        "items_cost": [
                            {
                                "item_id": 201001,
                                "count": 7
                            }
                        ],
                        "category": 1
                    }
                ],
                "op_time": 1485840313517,
                "msg_type": 102,
                "farm_id": "844424932415867"
            }
        ]
    }

喂食动物:
    http://api.welove520.com/v1/game/farm/animals/feed
参数:
    access_token=562949961343086-21410ab7bf7c154d0d&item_id=202001&sig=Q18smYGu5lBgL7vEc3CNUeDx5pw%3D&yards_info=%5B%7B%22yard_id%22%3A118817293%2C%22animals%22%3A%5B%7B%22id%22%3A6808593%2C%22last_interval%22%3A0%7D%2C%7B%22id%22%3A6808575%2C%22last_interval%22%3A0%7D%2C%7B%22id%22%3A6808561%2C%22last_interval%22%3A0%7D%5D%7D%5D
返回值:
    {
        "result": 1,
        "messages": [
            {
                "op_time": 1485840297373,
                "msg_type": 7,
                "yards": [
                    {
                        "animals": [
                            {
                                "animal_id": 6808593,
                                "last_feed_time": 1485840297371
                            }
                        ],
                        "id": 118817293
                    }
                ]
            },
            {
                "warehouses": [
                    {
                        "items_cost": [
                            {
                                "item_id": 203001,
                                "count": 3
                            }
                        ],
                        "category": 2
                    }
                ],
                "op_time": 1485840297373,
                "msg_type": 102,
                "farm_id": "844424932415867"
            }
        ]
    }

完成订单:
    http://api.welove520.com/v1/game/farm/order/accomplish
参数:
    access_token=562949961343086-21410ab7bf7c154d0d&by_rainbow_coin=0&order_id=30285736&sig=z9BSl%2BJSQa385JKfDpAyb3%2BpV1c%3D
返回值:
    {
        "result": 1,
        "messages": [
            {
                "special": 0,
                "time_left": 0,
                "voucher_item_id": 0,
                "op_time": 1485843157576,
                "msg_type": 47,
                "slot": 6,
                "exp": 163,
                "items": [
                    {
                        "item_id": 204033,
                        "count": 1
                    }
                ],
                "order_id": 37244643,
                "status": 0,
                "coin": 315,
                "buyer": 2
            },
            {
                "warehouses": [
                    {
                        "items_cost": [
                            {
                                "item_id": 201007,
                                "count": 1
                            }
                        ],
                        "category": 1
                    },
                    {
                        "items_cost": [
                            {
                                "item_id": 202003,
                                "count": 1
                            }
                        ],
                        "category": 2
                    }
                ],
                "op_time": 1485843157576,
                "msg_type": 102,
                "farm_id": "844424932415867"
            }
        ]
    }

订单完成奖金:
    http://api.welove520.com/v1/game/farm/order/reward
参数:
    access_token=562949961343086-21410ab7bf7c154d0d&order_id=30285736&sig=F8%2BQGxZUAUjgVV6hGK1TcVGn44I%3D
返回值:
    {
        "result": 1,
        "messages": [
            {
                "op_time": 1485843191724,
                "msg_type": 45
            },
            {
                "op_time": 1485843191724,
                "exp_inc": 237,
                "gold_inc": 313,
                "msg_type": 103,
                "farm_id": "844424932415867"
            }
        ]
    }

农场签到:
    http://api.welove520.com/v1/game/farm/signin
参数:
    access_token=****&app_key=****&ph=farm&sig=****
返回值:
    {
        "result": 1,
        "messages": [
            {
                "lover_records": [
                    {
                        "date": "2017-12-18",
                        "lasting_days": 0,
                        "invalid": 0,
                        "complete": 0
                    }
                ],
                "lasting_days": 115,
                "discount_card_left": 1,
                "max_cp_lasting_days": 86,
                "record": [
                    {
                        "date": "2017-12-18",
                        "lasting_days": 115,
                        "invalid": 0,
                        "complete": 1
                    }
                ],
                "today": "2017-12-18",
                "msg_type": 119,
                "lover_lasting_days": 86,
                "card_count": 5,
                "lover_card_count": 1
            }
        ]
    }
```
