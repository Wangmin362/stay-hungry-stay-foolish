fn main() {
    #[derive(PartialEq, Debug)]
    struct Point {
      x: f64,
      y: f64,
    }

    // 将Point实例装箱（放到堆内存）
    let box_point = Box::new(Point { x: 0.0, y: 0.0 });
    // 通过解引用操作符取出Point实例
    let unboxed_point: Point = *box_point;
    assert_eq!(unboxed_point, Point { x: 0.0, y: 0.0 });

    // 通过Deref技术，直接解引用，获取到内部的值
    assert_eq!(&box_point.y, 0.0);
    assert_eq!(&box_point.x, 0.0);
}