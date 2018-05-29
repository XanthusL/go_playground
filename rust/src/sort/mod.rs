pub trait Comparable {
    fn less(&self, target: &Self) -> i8;
}

pub fn bubble<T: Comparable>(src: &mut Vec<T>) {
//let ref mut tmp;
    for _ in 0..src.len() {
        for j in 0..src.len() - 1 {
            if (&src[j]).less(&src[j + 1]) < 0 {
//let (src[i],src[j]) = (src[j],src[i]);
// mem::swap(&mut src[j],&mut src[j+1]);
// tmp = &src[i];
// src[i] = &src[j];
// src[j] = *tmp;
                src.swap(j, j + 1);
            }
        }
    }
}


#[cfg(test)]
pub mod sort_test{
    fn new_person(age: i8) -> Person {
        return Person {
            name: "",
            age,
        };
    }
    struct Person {
        name: &'static str,
        age: i8,
    }

    impl super::Comparable for Person {
        fn less(&self, target: &Self) -> i8 {
            if self.age > target.age {
                return -1;
            } else if self.age < target.age {
                return 1;
            }
            return 0;
        }
    }

    #[test]
    fn bubble_test(){
        let mut people = vec![];
        people.push(new_person(8));
        people.push(new_person(9));
        people.push(new_person(6));
        super::bubble(&mut people);
        assert_eq!(people[0].age,6);
        assert_eq!(people[1].age,8);
        assert_eq!(people[2].age,9);
    }
}
